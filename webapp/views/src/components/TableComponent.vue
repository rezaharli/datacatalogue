<template>
  <v-data-table
    :items="items"
    :headers="displayedHeaders"
    :pagination.sync="opts.pagination"
    :total-items="opts.totalItems"
    :loading="opts.isLoading"
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
						:class="['column sortable text-xs-left', opts.pagination.descending ? 'desc' : 'asc', header.value === opts.pagination.sortBy ? 'active' : '']"
						@click="changeSort(header.value)"
					>
						<div class="th-wrapper" v-bind:style="{ width: header.width ? header.width : 'unset' }">
							<tableheader :opts="opts" :props="header" :which="'left'" :fromHeaderLoop="true" />
							<v-icon small>arrow_upward</v-icon>
						</div>
					</th>

					<th
						v-if="header.sortable == false"
						:key="header.text"
						:class="['column sortable text-xs-left', opts.pagination.descending ? 'desc' : 'asc', header.value === opts.pagination.sortBy ? 'active' : '']"
					>
						<div class="th-wrapper" v-bind:style="{ width: header.width ? header.width : 'unset' }">
							<tableheader :opts="opts" :props="header" :which="'left'" :fromHeaderLoop="true" />
						</div>
					</th>
				</template>
			</tr>
		</template>

    <v-progress-linear slot="progress" color="blue" indeterminate></v-progress-linear>

    <template slot="no-data">
      <v-alert :value="opts.isLoading" type="info">Please wait while data is loading</v-alert>

      <v-alert :value="!opts.isLoading" type="error">Sorry, nothing to display here</v-alert>
    </template>

    <template slot="items" slot-scope="props">
			<template v-for="header in displayedHeaders">
				<table-cell v-bind:key="header"
					showOn="hover" 
					:fulltext="props.item[header.value]" 
					:isTextLink="header.isLink" 
					:onClick="header.onClick" />
			</template>
    </template>
  </v-data-table>
</template>

<script>
import tableHeader from "@/components/TableHeader.vue";
import tableCell from "@/components/Tablecell.vue";

export default {
  name: "tableComponent",
  components: { tableHeader, tableCell },
  props: ["opts", "items"],
  data() {
    return {};
  },
  computed: {
		store() { return this.$store.state[this.headerStoreName]; },
		displayedHeaders() { return this.opts.headers.filter(v => v.display == true); }
  },
  watch: {},
  mounted() {},
  methods: {
    changeSort (column) {
      if (this.opts.pagination.sortBy === column) {
        this.opts.pagination.descending = !this.opts.pagination.descending;
      } else {
        this.opts.pagination.sortBy = column;
        this.opts.pagination.descending = false;
      }
    },
	}
};
</script>