<template>
  <v-data-table
    :items="items"
    :headers="displayedHeaders"
    :pagination.sync="pagination"
    :rows-per-page-items="pagination.rowsPerPageItems"
    :total-items="pagination.totalItems"
    :loading="loading"
    :must-sort="true"
    item-key="ID"
    class="card-content"
  >
    <!-- <template slot="headerCell" slot-scope="props">
      <table-header :storeName="storeName" :props="props" :which="'left'" />
    </template> -->

		<template slot="headers" slot-scope="props">
			<tr>
				<!-- <th>
					<v-checkbox :input-value="props.all" :indeterminate="props.indeterminate" primary hide-details @click.stop="toggleAll"></v-checkbox>
				</th> -->

				<template v-for="header in props.headers">
					<th
						v-if="header.sortable == true"
						:key="header.text"
						:class="['column sortable text-xs-left', pagination.descending ? 'desc' : 'asc', header.value === pagination.sortBy ? 'active' : '']"
						@click="changeSort(header.value)"
					>
						<div class="th-wrapper" v-bind:style="{ width: header.width ? header.width : 'unset' }">
							<table-header :opts="opts" :props="header" :which="'left'" :fromHeaderLoop="true" />
							<v-icon small>arrow_upward</v-icon>
						</div>
					</th>

					<th
						v-if="header.sortable == false"
						:key="header.text"
						:class="['column sortable text-xs-left', pagination.descending ? 'desc' : 'asc', header.value === pagination.sortBy ? 'active' : '']"
					>
						<div class="th-wrapper" v-bind:style="{ width: header.width ? header.width : 'unset' }">
							<table-header :opts="opts" :props="header" :which="'left'" :fromHeaderLoop="true" />
						</div>
					</th>
				</template>
			</tr>
		</template>

    <v-progress-linear slot="progress" color="blue" indeterminate></v-progress-linear>

    <template slot="no-data">
      <v-alert :value="loading" type="info">Please wait while data is loading</v-alert>

      <v-alert :value="!loading" type="error">Sorry, nothing to display here</v-alert>
    </template>

    <template slot="items" slot-scope="props">
      <tr :class="{even: props.index % 2, odd: !(props.index % 2)}">
        <table-cell v-bind:key="header.text+i" v-for="(header, i) in displayedHeaders"
          showOn="hover" 
          :fulltext="props.item[header.value]" 
          :isTextLink="header.isLink" 
          :onClick="header.onClick" />
      </tr>
    </template>
  </v-data-table>
</template>

<script>
import tableHeader from "@/components/TableHeader.vue";
import tableCell from "@/components/Tablecell.vue";

export default {
  name: "tableComponent",
  components: { tableHeader, tableCell },
  props: ["opts"],
  data() {
    return {
      loading: false
    };
  },
  computed: {
    displayedHeaders() { return this.opts.headers.filter(v => v.display == true); },
    tableState(){ return this.$store.getters['tableModule/allState']; },
    items: {
      get: function () {
        return this.$store.getters['tableModule/items'];
      },
      set: function (value) {
        this.$store.commit('tableModule/setItems', value);
      }
    },
    pagination: {
      get: function () {
        return this.$store.getters['tableModule/pagination'];
      },
      set: function (value) {
        this.$store.commit('tableModule/setPagination', value);
      }
    },
    filters: {
      get: function () {
        return this.$store.getters['tableModule/filters'];
      },
      set: function (value) {
        this.$store.commit('tableModule/setFilters', value);
      }
    },
  },
  watch: {
    pagination: {
      handler () {
        this.getTableData();
      },
      deep: true
    },
    items: {
      handler () {
        console.log("a");
        
      },
      deep: true
    },
  },
  mounted() {},
  methods: {
    changeSort(column) {
      if (this.pagination.sortBy === column) {
        this.pagination.descending = !this.pagination.descending;
      } else {
        this.pagination.sortBy = column;
        this.pagination.descending = false;
      }
    },
    getTableData(){
      Object.keys(this.filters).map(function(key, index) {
        this.filters[key] = (typeof(this.filters[key]) == "object") ? this.filters[key] : (this.filters[key] ? this.filters[key].toString() : "");
      });

      this.tableState.payload = {
        Filters: this.filters,
        Pagination: this.pagination
      }

      this.opts.data.read(this.tableState);
      
      return new Promise((resolve, reject) => {
      }).then(() => {
        console.log(o.items);
      });
    },
    setTableData(){
      this.loading = true;
      
      this.getTableData().then(() => {
        this.$store.dispatch('tableModule/setData', res).then(() => {
          this.loading = false;
        });
      })
    }
	}
};
</script>