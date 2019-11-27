<template>
  <div class="wrapper-showmore d-inline-block">
    <span v-if="showOn == 'click'" @click="toggleText();">{{ isTruncated ? truncateCells(fulltext) : fulltext }}</span>

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
  mounted() {},
  methods: {
    toggleText() {
      this.isTruncated = !this.isTruncated;
    },
    truncateCells(text) {
      var n = 15;

      if( ! text) return "";

      return text.length > n ? text.slice(0, n) + "..." : text;
    },
    onClick: function (e, isTruncated) {
      var clickedElem = e.currentTarget;
      var tdElem = $(clickedElem).closest("td");
      var keberapa = tdElem.index();
      setTimeout(function(){
        var tdWidth = tdElem.outerWidth();
        var tdWidthMax = Math.max.apply(Math, tdElem.closest("table > tbody > tr > td").eq(keberapa).map(function(){ return $(this).outerWidth(); }).get());
        var wrapperShowMoreWidthMax = Math.max.apply(Math, tdElem.closest("table > tbody > tr > td").eq(keberapa).children(".wrapper-showmore").map(function(){ return $(this).outerWidth(); }).get());
        var thWidth = tdElem.closest("table").children("thead").children("tr").children("th").eq(keberapa).attr("data-width-ori");
        
        if(isTruncated){
          if(wrapperShowMoreWidthMax>thWidth){
            var tdWidthUsed = tdWidthMax;
          }else{
            var tdWidthUsed = thWidth;
          }
        }else{
          var tdWidthUsed = tdWidthMax;
        }
        tdElem.closest("table").children("tbody").children("tr").each(function () {
          $(this).children("td").eq(keberapa).css({"width": tdWidthUsed+"px"});
        });
        tdElem.closest("table").children("thead").children("tr").children("th").eq(keberapa).css({"min-width": tdWidthUsed+"px"});
      }, 10);
    }
  }
};
</script>