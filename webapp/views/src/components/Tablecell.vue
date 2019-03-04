<template>
  <div>
    <span v-if="isklik" @click="toggleText">{{ isTruncated ? truncateCells(fulltext) : fulltext }}</span>
    <span v-if="!isklik">{{ isTruncated ? truncateCells(fulltext) : fulltext }}</span>
  </div>
</template>

<script>
  export default {
    name: "tablecell",
    props: ["fulltext", "isklik"],
    data () {
      return {
        isTruncated: true,
      }
    },
    methods: {
      toggleText(){
        this.isTruncated = !this.isTruncated
      },
      truncateCells(text){
        var n = 15;
        return text.length > n ? text.slice(0, n) + "..." : text;
      }
    }
  }
</script>