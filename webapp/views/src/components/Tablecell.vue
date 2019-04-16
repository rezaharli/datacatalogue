<template>
  <div>
    <span v-if="showOn == 'click'" @click="toggleText">{{ isTruncated ? truncateCells(fulltext) : fulltext }}</span>

    <p
      style="margin-bottom: 0px !important;"
      v-if="showOn == 'hover'"
      @mouseenter="toggleText"
      @mouseleave="toggleText"
    >{{ isTruncated ? truncateCells(fulltext) : fulltext }}</p>
  </div>
</template>

<script>
export default {
  name: "tablecell",
  props: ["fulltext", "showOn"],
  data() {
    return {
      isTruncated: true
    };
  },
  methods: {
    toggleText() {
      this.isTruncated = !this.isTruncated;
    },
    truncateCells(text) {
      var n = 15;

      if( ! text) return "";

      return text.length > n ? text.slice(0, n) + "..." : text;
    }
  }
};
</script>