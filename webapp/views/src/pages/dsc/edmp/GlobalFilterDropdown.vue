<template>
  <v-tooltip bottom>
    <template slot="activator" slot-scope="{ on }">
      <v-layout row wrap v-on="on">
        <v-flex d-flex xs12 style="flex-direction: column;">
          <v-autocomplete v-on="on" single-line multiple box clearable
            v-model="ddValues"
            :items="mappedItems"
            :search-input.sync="searchInput"
            :label="label"
            :disabled="disabled"
            :readonly="disabled"
            :loading="disabled"
            item-text="name"
          >
            <template slot="selection" slot-scope="{ item, index }">
              <v-chip v-if="index === 0" :class="ddValues.length == 1 ? 'full' : 'small'">
                <span>{{ item.name }}</span>
              </v-chip>

              <span v-if="index === 1" class="grey--text caption">({{ ddValues.length - 1 }}+)</span>
            </template>
          </v-autocomplete>
        </v-flex>
      </v-layout>
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
  props: ["value", "items", "label", "disabled"],
  data() {
    return {
      searchInput: ""
    };
  },
  computed: {
    mappedItems() {
      return this.items.map(v => {
        return { name: v, disabled: this.disabled };
      });
    },
    ddValues: {
      get: function() {
        return this.value;
      },
      set: function(newValue) {
        this.$emit("input", newValue);
      }
    },
  },
  methods: {}
};
</script>
