<style>
  @import '../../assets/styles/dashboard.css';
</style>

<template>
    <v-content>
        <v-container fluid>
            <v-layout row wrap>
                <v-flex xs6>
                    <v-data-table
                        :headers="headers"
                        :items="desserts"
                        :loading="true"
                        class="elevation-1">
                        <v-progress-linear slot="progress" color="blue" indeterminate></v-progress-linear>
                        <template slot="items" slot-scope="props">
                            <td>{{ props.item.name }}</td>
                            <td>{{ props.item.calories }}</td>
                            <td>{{ props.item.fat }}</td>
                        </template>
                    </v-data-table>
                </v-flex>
            </v-layout>
        </v-container>
    </v-content>
</template>

<script>
export default {
    data () {
      return {
        headers: [
          { text: 'System Name (4)', align: 'left', sortable: false, value: 'name' },
          { text: 'ITAM ID (4)', align: 'left', sortable: false, value: 'calories' },
          { text: 'Dataset Custodian', align: 'left', sortable: false, value: 'fat' },
        ],
        desserts: [
          {
            name: 'Frozen Yogurt',
            calories: 159,
            fat: 6.0,
            carbs: 24,
          },
          {
            name: 'Ice cream sandwich',
            calories: 237,
            fat: 9.0,
            carbs: 37,
          },
          {
            name: 'Eclair',
            calories: 262,
            fat: 16.0,
            carbs: 23,
          },
          {
            name: 'Cupcake',
            calories: 305,
            fat: 3.7,
            carbs: 67,
          },
          {
            name: 'Gingerbread',
            calories: 356,
            fat: 16.0,
            carbs: 49,
          },
          {
            name: 'Jelly bean',
            calories: 375,
            fat: 0.0,
            carbs: 94,
          },
          {
            name: 'Lollipop',
            calories: 392,
            fat: 0.2,
            carbs: 98,
          },
          {
            name: 'Honeycomb',
            calories: 408,
            fat: 3.2,
            carbs: 87,
          },
          {
            name: 'Donut',
            calories: 452,
            fat: 25.0,
            carbs: 51,
          },
          {
            name: 'KitKat',
            calories: 518,
            fat: 26.0,
            carbs: 65,
          }
        ]
      }
    }
}
</script>
