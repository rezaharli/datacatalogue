<template>
  <p slot="activator" class="col-form-label">
    <span v-html="truncatedText"></span>

    <v-dialog v-model="dialog" width="500" style="margin-top: -4px;" v-if="isTruncated">
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
      isTruncated: false
    };
  },
  computed: {
    truncatedText () {
      if( ! this.fulltext) return "";
      
      return this.truncateDiz(this.fulltext)
    }
  },
  methods: {
    truncateDiz(text) {
      var n = 100;

      if (text.length > n){
        this.isTruncated = true;
        return text.slice(0, n) + "... ";
      } else {
        this.isTruncated = false;
        return text;
      }
    }
  }
};
</script>