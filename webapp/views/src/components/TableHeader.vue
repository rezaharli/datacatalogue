<template>
  <div ref="widthAcuan">
    {{ props.header.text }} ({{ store[which].source[0] ? store[which].source[0]["COUNT_" + props.header.value.split(".").reverse()[0]] : 0 }})

    <b-dropdown no-caret variant="link" class="header-filter-icon" ref="columnFilter">
      <template slot="button-content">
        <!-- <i class="fa fa-filter text-muted"></i> -->
        <v-icon small>filter_list</v-icon>
      </template>

      <b-dropdown-header>
        <b-form-input
          type="text"
          placeholder="Filter"
          v-model="store.filters[which][props.header.value.split('.').reverse()[0]]"
          @keyup.native="keyupAction"
          @change="filterKeyup(props.header)"
        ></b-form-input>
      </b-dropdown-header>

      <b-dropdown-item
        v-for="item in distinctData(props.header.value, store[which].source)"
        :key="item"
        @click="filterClick(props.header, item)"
      >{{ item }}</b-dropdown-item>
    </b-dropdown>
  </div>
</template>

<script>
export default {
  name: "tableHeader",
  props: ["storeName", "props", "which"],
  data() {
    return {};
  },
  computed: {
      store () { return this.$store.state[this.storeName].all },
  },
  mounted (){
    setTimeout(
      () => this.store[this.which].colWidth[this.props.header.value.split('.').reverse()[0]] = this.$refs.widthAcuan.parentNode.offsetWidth, 
    100)
  },
  methods: {
    getLeftTable () {
        return this.$store.dispatch(`${this.storeName}/getLeftTable`)
    },
    getRightTable (id) {
        this.$store.dispatch(`${this.storeName}/getRightTable`, id)
    },
    distinctData (col, datax) {
      var cols = col.split(".")

      if(cols.length > 1){
        var a = datax;

        cols.forEach((c, i) => {
          a = this._.flattenDeep(this._.map(this._.sortBy(a, c), c));
        });

        return this._.uniq(a).filter(Boolean);
      }
      
      return this._.uniq(
        this._.map(this._.sortBy(datax, col), col)
      ).filter(Boolean);
    },
    keyupAction(e){
      if(e.key == "Enter") this.$refs.columnFilter.hide(true)
    },
    filterKeyup (keyModel) {
      if(this.which == "left") this.getLeftTable()
      else this.getRightTable(this.$route.params.system)
    },
    filterClick (keyModel, val) {
      this.store.filters[this.which][keyModel.value.split('.').reverse()[0]] = val;

      if(this.which == "left") this.getLeftTable()
      else this.getRightTable(this.$route.params.system)
    },
  }
};
</script>