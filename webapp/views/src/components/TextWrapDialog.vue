<template>
  <p slot="activator" class="col-form-label">
    <span v-html="truncatedText"></span>
    <v-dialog v-model="dialog" width="500" style="margin-top: -4px;" v-if="truncatedText != fulltext">
      <b-link slot="activator" style="margin-top: -4px;">[more]</b-link>

      <v-card>
        <v-card-text v-html="fulltext"></v-card-text>
      </v-card>
    </v-dialog>
  </p>
</template>

<script>
export default {
  name: "textWrapDialog",
  props: ["fulltext"],
  data() {
    return {
      dialog: false,
    };
  },
  computed: {
    truncatedText () {
      return this.truncateDiz(this.fulltext)
    }
  },
  methods: {
    truncateDiz(text) {
      var n = 100;
      return text.length > n ? text.slice(0, n) + " " : text;
    }
  }
};
</script>