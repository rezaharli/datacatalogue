<style>
.v-menu__content {
  cursor: default;
  padding: 0.5rem 0;
}
</style>


<template>
  <div ref="widthAcuan" class="table-header-wrapper">
    <span>{{ fixedProps.header.text }} {{ count }}</span>

    <v-menu
      bottom
      transition="slide-y-transition"
      class="dropdown-button-wrapper"
      v-if="fixedProps.header.filterable"
      v-model="menu"
      :close-on-content-click="false"
    >
      <template slot="activator" slot-scope="{ on }">
        <button class="dropdown-button-wrapper btn btn-link">
          <v-icon
            small
            class="icon-filter"
            v-on="on"
            v-bind:class="{ 'icon-active': tableStore.filters[which][fixedProps.header.value.split('.').reverse()[0]] }"
          >filter_list</v-icon>
        </button>
      </template>

      <b-dropdown-header>
        <b-form-input
          type="text"
          placeholder="Filter"
          v-model="tableStore.filters[which][fixedProps.header.value.split('.').reverse()[0]]"
          @keyup.native="keyupAction"
        ></b-form-input>
      </b-dropdown-header>

      <b-dropdown-divider />

      <div class="dropdown-wrapper">
        <page-loader v-if="isLoading" />

        <b-dropdown-item
          v-for="item in dropdownData"
          :key="item"
          @click="filterClick(fixedProps.header, item)"
        >{{ item }}</b-dropdown-item>
      </div>

      <b-dropdown-divider />

      <b-row
        class="justify-content-center"
        v-bind:class="{ 'd-none': !(tableStore.filters[which][fixedProps.header.value.split('.').reverse()[0]]) }"
      >
        <b-col cols="auto">
          <a
            class="text-danger mx-4 my-1"
            @click="resetFilterColumn(which, fixedProps.header.value.split('.').reverse()[0])"
          >
            <i class="fa fa-trash"></i> Clear
          </a>
        </b-col>
      </b-row>
    </v-menu>
  </div>
</template>

<script>
import pageLoader from "../PageLoader.vue";

export default {
  name: "tableHeader",
  components: {
    pageLoader
  },
  props: ["storeName", "props", "which", "fromHeaderLoop"],
  data() {
    return {
      headerStoreName: "header",
      edmpStoreName: "edmp",
      filterProcessTimeout: null,
      sticky: 0,
      menu: false,
      isLoading: false,
      dropdownData: []
    };
  },
  computed: {
    store() {
      return this.$store.state[this.headerStoreName];
    },
    edmpStore() {
      return this.$store.state.edmp.all;
    },
    tableStore() {
      return this.$store.state[this.storeName].all;
    },
    count() {
      if (!this.fixedProps.header.displayCount) return "";
      if (!this.tableStore[this.which].source[0]) return "(0)";

      return (
        "(" +
        this.tableStore[this.which].source[0][
          "COUNT_" + this.fixedProps.header.value.split(".").reverse()[0]
        ] +
        ")"
      );
    },
    distinctData() {
      var headerValueField = this.fixedProps.header.value;
      var datax = this.tableStore[this.which].source;

      var cols = headerValueField.split(".");

      if (cols.length > 1) {
        var a = datax;

        cols.forEach((c, i) => {
          a = this._.flattenDeep(this._.map(this._.sortBy(a, c), c));
        });

        return this._.uniq(a).map(v => (v.toString() ? v.toString() : "NA"));
      }

      return this._.uniq(
        this._.map(
          this._.sortBy(datax, headerValueField),
          headerValueField
        ).map(v => (v.toString() ? v.toString() : "NA"))
      );
    },
    fixedProps() {
      if (this.fromHeaderLoop == true) {
        var props = {};
        props.header = this.props;

        return props;
      } else {
        return this.props;
      }
    }
  },
  watch: {
    menu(val, oldVal) {
      if (val) {
        this.getOpts();
      } else {
        this.dropdownData = [];
      }
    }
  },
  mounted() {
    // this.makeTableAlertFull();
    setTimeout(() => {
      var $dropdownWrapper = $(this.$refs.widthAcuan)
        .closest("th.column.sortable")
        .find(".dropdown-button-wrapper");
      if ($dropdownWrapper[0]) {
        $dropdownWrapper.on("click", function(e) {
          e.preventDefault();
          e.stopPropagation();
        });
      }
    }, 100);
  },
  methods: {
    getOpts() {
      this.isLoading = true;

      var param = this.tableStore.param;
      param.Filename = this.tableStore.filename;
      param.Queryname = this.tableStore.queryname;
      param.FieldName = this.fixedProps.header.value;
      param.GlobalFilters = this.edmpStore.globalFilters;

      param.ColumnFilters = this.tableStore.filters[this.which][
        this.fixedProps.header.value.split(".").reverse()[0]
      ];

      return this.$store
        .dispatch(`${this.headerStoreName}/getOpts`, param)
        .then(() => {
          this.dropdownData = this._.uniq(
            this.store.datas.map(v =>
              v.toString().trim() ? v.toString() : "NA"
            )
          );

          this.isLoading = false;
        });
    },
    getLeftTable() {
      return this.$store.dispatch(`${this.storeName}/getLeftTable`);
    },
    getRightTable(id) {
      this.$store.dispatch(`${this.storeName}/getRightTable`, id);
    },
    keyupAction(e) {
      if (this.tableStore.filters[this.which].filterTypes){
        delete this.tableStore.filters[this.which].filterTypes[
          this.fixedProps.header.value.split(".").reverse()[0]
        ];
      }

      setTimeout(this.getOpts, 1);

      var self = this;

      // manually adding space only when sortable is true
      if (e.key == " ") {
        if (this.fixedProps.header.sortable) {
          var model = this.tableStore.filters[this.which][
            this.fixedProps.header.value.split(".").reverse()[0]
          ];
          model = model ? model + " " : " ";

          this.tableStore.filters[this.which][
            this.fixedProps.header.value.split(".").reverse()[0]
          ] = model;
        }
      }

      if (e.key == "Enter") {
        this.menu = false;

        if (self.filterProcessTimeout != null)
          clearTimeout(self.filterProcessTimeout);

        self.filterProcessTimeout = setTimeout(self.filterProcess, 500);
      }
    },
    filterClick(keyModel, val) {
      this.tableStore.filters[this.which][
        keyModel.value.split(".").reverse()[0]
      ] = val;

      this.tableStore.filters[this.which].filterTypes = {};
      this.tableStore.filters[this.which].filterTypes[
        keyModel.value.split(".").reverse()[0]
      ] = "eq";

      this.filterProcess();
    },
    filterProcess() {
      if (this.which == "left") this.getLeftTable();
      else this.getRightTable(this.$route.params.system);
    },
    resetFilterColumn(which, fieldName) {
      this.tableStore.filters[which][fieldName] = "";
      this.tableStore.filters[which].filterTypes = {};
      this.filterProcess();
      this.menu = false;
    },
    // onScrollListener(e) {
    //   // metode mengawang alias menambah class sticky ke current elemen
    //   if ($(window).scrollTop() > this.sticky) {
    //     // && tableBody.height() > window.innerHeight
    //     $("table.v-table.v-datatable thead").addClass("sticky");
    //     $("table.v-table.v-datatable thead").css({
    //       top: $(".v-toolbar").height()
    //     });
    //     // $('.v-datatable__actions').css({'position': 'fixed', 'bottom': 0});
    //   } else {
    //     $("table.v-table.v-datatable thead").removeClass("sticky");
    //     $("table.v-table.v-datatable thead").css({ top: "unset" });
    //     // $('.v-datatable__actions').css({'position': 'unset', 'bottom': 'unset'});
    //   }
    // },

    // makeTableHeaderFixed() {
    //   const theads = document.querySelectorAll(
    //     "table.v-table.v-datatable thead"
    //   );
    //   theads.forEach(thead => {
    //     thead.querySelectorAll("tr > th").forEach(th => {
    //       th.style.width = th.offsetWidth + "px";
    //     });
    //   });

    //   this.sticky = $("table.v-table thead").offset().top;
    //   $(window).scroll(this.onScrollListener);
    // },

    // makeTableAlertFull() {
    //   var elemAlert = $("table.v-table > tbody > tr > td >  .v-alert");
    //   var theadWidth =
    //     elemAlert
    //       .closest("table.v-table")
    //       .find("thead > tr:first")
    //       .width() - 50;
    //   elemAlert.parent("td").width(theadWidth);
    // }
  }
};
</script>