<style>
  .dropdown-menu {
    cursor: default;
  }
</style>


<template>
  <div ref="widthAcuan" class="table-header-wrapper">
    <span>{{ props.header.text }} {{ count }}</span>

    <b-dropdown no-caret variant="link" class="dropdown-button-wrapper" ref="columnFilter" v-if="props.header.filterable">
      <template slot="button-content">
        <v-icon small v-bind:class="{'icon-active' : store.filters[which][props.header.value.split('.').reverse()[0]] }" class="icon-filter">filter_list</v-icon>
      </template>

      <b-dropdown-header>
        <b-form-input
          type="text"
          placeholder="Filter"
          v-model="store.filters[which][props.header.value.split('.').reverse()[0]]"
          @keyup.native="keyupAction"
        ></b-form-input>
      </b-dropdown-header>

      <b-dropdown-divider/>

      <div class="dropdown-wrapper">
        <b-dropdown-item
          v-for="item in distinctData"
          :key="item"
          @click="filterClick(props.header, item)"
        >{{ item }}</b-dropdown-item>
      </div>

      <b-dropdown-divider/>

      <b-row class="justify-content-center" v-bind:class="{'d-none' : !(store.filters[which][props.header.value.split('.').reverse()[0]]) }">
        <b-col cols="auto">
          <a class="text-danger mx-4 my-1" @click="resetFilterColumn(which, props.header.value.split('.').reverse()[0])">
            <i class="fa fa-trash"></i> Clear
          </a>
        </b-col>
      </b-row>

    </b-dropdown>
  </div>
</template>

<script>
export default {
  name: "tableHeader",
  props: ["storeName", "props", "which"],
  data() {
    return {
      filterProcessTimeout: null,
      sticky: 0
    };
  },
  computed: {
    store () { return this.$store.state[this.storeName].all },
    count () {
      if( ! this.props.header.displayCount) return "";
      if( ! this.store[this.which].source[0]) return "(0)";

      return ("(" + this.store[this.which].source[0]["COUNT_" + this.props.header.value.split(".").reverse()[0]] + ")");
    },
    distinctData () {
      var headerValueField = this.props.header.value;
      var datax = this.store[this.which].source;

      var cols = headerValueField.split(".")

      if(cols.length > 1){
        var a = datax;

        cols.forEach((c, i) => {
          a = this._.flattenDeep(this._.map(this._.sortBy(a, c), c));
        });

        var ret = this._.uniq(a).map(v => v.toString().trim()).filter(Boolean);

        return ret;
      }
      
      return this._.uniq(
        this._.map(this._.sortBy(datax, headerValueField), headerValueField).map(v => v.toString().trim())
      ).filter(Boolean);
    },
  },
  mounted (){
    this.makeTableAlertFull();
    setTimeout(() => {
      var $dropdownWrapper = $(this.$refs.widthAcuan).closest("th.column.sortable").find(".dropdown-menu");
      if($dropdownWrapper[0]){
        $dropdownWrapper.on("click", function(e) {
          e.preventDefault();
          e.stopPropagation();
        });
      }

      // this.store[this.which].colWidth[this.props.header.value.split('.').reverse()[0]] = this.$refs.widthAcuan.parentNode.offsetWidth;
      // console.log(this.$refs.widthAcuan.parentNode, '----' ,this.$refs.widthAcuan.parentNode.offsetWidth);
    }, 100);
    setTimeout(() => {
      // this.makeTableHeaderFixed();
    }, 200);
  },
  methods: {
    getLeftTable () {
        return this.$store.dispatch(`${this.storeName}/getLeftTable`)
    },
    getRightTable (id) {
        this.$store.dispatch(`${this.storeName}/getRightTable`, id)
    },
    keyupAction(e){
      var self = this;

      // manually adding space only when sortable is true
      if(e.key == " ") {
        if(this.props.header.sortable){
          var model = this.store.filters[this.which][this.props.header.value.split('.').reverse()[0]];
          model = model ? model + " " : " ";

          this.store.filters[this.which][this.props.header.value.split('.').reverse()[0]] = model;

        }
      }

      if(self.filterProcessTimeout != null) clearTimeout(self.filterProcessTimeout);

      self.filterProcessTimeout = setTimeout(self.filterProcess, 500);

      if(e.key == "Enter") self.$refs.columnFilter.hide(true)
    },
    filterClick (keyModel, val) {
      this.store.filters[this.which][keyModel.value.split('.').reverse()[0]] = val;
      
      this.store.filters[this.which].filterTypes = {};
      this.store.filters[this.which].filterTypes[keyModel.value.split('.').reverse()[0]] = "eq"

      this.filterProcess();
    },
    filterProcess (){
      if(this.which == "left") this.getLeftTable()
      else this.getRightTable(this.$route.params.system)
    },
    resetFilterColumn (which, fieldName) {
      // console.log(this.store.filters[this.which]);
      // console.log("-------------------");
      // console.log("fieldName : ", fieldName);
      // console.log("-------------------");
      // console.log(this.store.filters[which][fieldName]);
      
      this.store.filters[which][fieldName] = "";
      this.store.filters[which].filterTypes = {};
      this.filterProcess ();
      this.$refs.columnFilter.hide(true);
    },
    onScrollListener(e){
        // metode mengawang alias menambah class sticky ke current elemen 
        if ($(window).scrollTop() > this.sticky) {
            // && tableBody.height() > window.innerHeight 
            $('table.v-table.v-datatable thead').addClass("sticky");
            $('table.v-table.v-datatable thead').css({'top': $('.v-toolbar').height()});
            // $('.v-datatable__actions').css({'position': 'fixed', 'bottom': 0});
        } else {
            $('table.v-table.v-datatable thead').removeClass("sticky");
            $('table.v-table.v-datatable thead').css({'top': 'unset'});
            // $('.v-datatable__actions').css({'position': 'unset', 'bottom': 'unset'});
        }

        // // pake clone jes
        // if ($(window).scrollTop() > this.sticky) {
        //     // && tableBody.height() > window.innerHeight 
        //     $('table.v-table.v-datatable thead.sticky').show();
        //     $('table.v-table.v-datatable thead.sticky').css({'top': $('.v-toolbar').height()});
        // } else {
        //     $('table.v-table.v-datatable thead.sticky').hide();
        //     $('table.v-table.v-datatable thead.sticky').css({'top': 'unset'});
        // }
    },

    makeTableHeaderFixed(){
        const theads = document.querySelectorAll('table.v-table.v-datatable thead');
        theads.forEach((thead) => {
          thead.querySelectorAll("tr > th").forEach((th) => {
            th.style.width = th.offsetWidth+'px';
          });

          // let clone = thead.cloneNode( true );
          // clone.setAttribute( 'class', 'sticky');
          // thead.closest('table.v-table.v-datatable').appendChild( clone );
          // clone.style.display = "none";
        });

        this.sticky = $('table.v-table thead').offset().top;    
        $(window).scroll(this.onScrollListener);
    },

    makeTableAlertFull(){
      var elemAlert = $('table.v-table > tbody > tr > td >  .v-alert');
      var theadWidth = elemAlert.closest('table.v-table').find('thead > tr:first').width() - 50;
      elemAlert.parent('td').width(theadWidth);
    },
  }
};
</script>