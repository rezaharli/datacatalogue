<template>
  <v-data-table
    :items="items"
    :headers="opts.headers.filter(v => v.display == true)"
    :pagination.sync="opts.pagination"
    :total-items="opts.totalItems"
    :loading="opts.isLoading"
    :must-sort="true"
    item-key="ID"
    class="card-content"
  >
    <template slot="headerCell" slot-scope="props">
      <table-header :storeName="storeName" :props="props" :which="'left'" />
    </template>

    <v-progress-linear slot="progress" color="blue" indeterminate></v-progress-linear>

    <template slot="no-data">
      <v-alert :value="opts.isLoading" type="info">Please wait while data is loading</v-alert>

      <v-alert :value="!opts.isLoading" type="error">Sorry, nothing to display here</v-alert>
    </template>

    <template slot="items" slot-scope="props">
      <td v-bind:style="{ width: opts.headers.width['SYSTEM_NAME'] + 'px' }">
        <b-link @click="showRightTable(props.item)">
          <table-cell :fulltext="props.item.SYSTEM_NAME" showOn="hover"></table-cell>
        </b-link>
      </td>

      <td v-bind:style="{ width: opts.headers.width['ITAM_ID'] + 'px' }">
        <b-link @click.stop="toggleDscDrawer(props.item)">
          <table-cell
            :fulltext="(_.uniq(_.map(props.item.Custodians, 'ITAM_ID').filter(Boolean)).join(', '))"
            showOn="hover"
          ></table-cell>
        </b-link>
      </td>
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
    store() {
      return this.$store.state[this.headerStoreName];
    }
  },
  watch: {},
  mounted() {},
  methods: {}
};
</script>