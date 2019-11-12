<template>
  <v-tooltip bottom>
    <template slot="activator" slot-scope="{ on }">
      <div v-on="on">
        <v-autocomplete single-line multiple box clearable
          v-model="ddValues"
          :items="items"
          :search-input.sync="searchInput"
          :label="label"
        >
          <template slot="selection" slot-scope="{ item, index }">
            <v-chip v-if="index === 0" :class="ddValues.length == 1 ? 'full' : 'small'">
              <span>{{ item }}</span>
            </v-chip>

            <span v-if="index === 1" class="grey--text caption">({{ ddValues.length - 1 }}+)</span>
          </template>
        </v-autocomplete>
      </div>
    </template>

    <div>
      <span>{{ label }}</span>

      <ul v-if="ddValues.length > 0">
        <li :key="val" v-for="val in ddValues">{{ val }}</li>
      </ul>
    </div>
  </v-tooltip>
</template>

<script>
export default {
  name: "globalFilterDropdown",
  props: ["value", "items", "label"],
  data() {
    return {
      searchInput: ""
    };
  },
  computed: {
    ddValues: {
      get: function() {
        return this.value;
      },
      set: function(newValue) {
        this.$emit("input", newValue);
      }
    }
  },
  methods: {}
};
</script>
