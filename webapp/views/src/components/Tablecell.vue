<template>
  <td>
    <b-link v-if="isTextLink" @click="showRightTable(props.item)">
      <div class="wrapper-showmore d-inline-block">
        <span v-if="showOn == 'click'" @click="toggleText();">{{ isTruncated ? truncateCells(fulltext) : fulltext }}</span>

        <p v-if="showOn == 'hover'"
          style="margin-bottom: 0px !important;"
          @mouseenter="toggleText"
          @mouseleave="toggleText"
        >{{ isTruncated ? truncateCells(fulltext) : fulltext }}</p>
      </div>
    </b-link>
  </td>
</template>

<script>
export default {
  name: "tablecell",
  props: ["fulltext", "showOn", "isTextLink"],
  data() {
    return {
      isTruncated: true
    };
  },
  mounted() {},
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