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
            v-bind:class="{ 'icon-active': tableStore.filters[fixedProps.header.value.split('.').reverse()[0]] }"
          >filter_list</v-icon>
        </button>
      </template>

      <b-dropdown-header>
        <b-form-input
          type="text"
          placeholder="Filter"
          v-model="tableStore.filters[fixedProps.header.value.split('.').reverse()[0]]"
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
          v-bind:class="{ 'd-none': !(tableStore.filters[fixedProps.header.value.split('.').reverse()[0]]) }"
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
import pageLoader from "./PageLoader.vue";

export default {
  name: "tableHeader",
  components: {
    pageLoader
  },
  props: ["opts", "props", "which", "fromHeaderLoop"],
  data() {
    return {
      headerStoreName: "header",
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
    tableStore() {
      return this.$store.state.tableModule;
    },
    fixedProps() {
      if (this.fromHeaderLoop == true) {
        var props = {};
        props.header = this.props;

        return props;
      } else {
        return this.props;
      }
    },
    count() {
      if (!this.fixedProps.header.displayCount) return "";
      if (!this.tableStore.items[0]) return "(0)";

      return (
        "(" +
        this.tableStore.items[0][
          "COUNT_" + this.fixedProps.header.value.split(".").reverse()[0]
        ] +
        ")"
      );

      return "(0)";
    },
  },
  watch: {
    menu(val, oldVal) {
      setTimeout(() => {
        var filterValue = this.tableStore.filters[this.fixedProps.header.value.split(".").reverse()[0]];
        if(filterValue == undefined || filterValue == ""){
          if (val) {
            this.getOpts();
          } else {
            this.dropdownData = [];
          }
        }
      }, 0);
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

      param.ColumnFilters = this.tableStore.filters[
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
    keyupAction(e) {
      if (this.tableStore.filters.filterTypes){
        delete this.tableStore.filters.filterTypes[
          this.fixedProps.header.value.split(".").reverse()[0]
        ];
      }

      setTimeout(this.getOpts, 1);

      var self = this;

      // manually adding space only when sortable is true
      if (e.key == " ") {
        if (this.fixedProps.header.sortable) {
          var model = this.tableStore.filters[
            this.fixedProps.header.value.split(".").reverse()[0]
          ];
          model = model ? model + " " : " ";

          this.tableStore.filters[
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
      this.menu = false;

      var fieldName = keyModel.value.split(".").reverse()[0];
      this.tableStore.filters[fieldName] = val;

      this.tableStore.filters.filterTypes = {};
      this.tableStore.filters.filterTypes[fieldName] = "eq";

      this.dropdownData = [val]

      this.filterProcess();
    },
    filterProcess() {
      this.opts.fetch();
    },
    resetFilterColumn(which, fieldName) {
      this.tableStore.filters[fieldName] = "";
      this.tableStore.filters.filterTypes = {};

      this.filterProcess();

      this.menu = false;
    },
  }
};
</script>