<template>
  <div>
  <div class="section" style="backgroundColor:#111">
    <div class="subtitle has-text-white">
      Shortlinks | <a href="https://jlyon.org" style="color:white; textDecoration:none;">jlyon.org</a>
    </div>
  </div>
  <div class = "section">
    {{links}}
  </div>
  </div>
</template>

<script>
export default {
  name: 'Main',
  data () {
    return {
      links: []
    }
  },
  mounted: function(){
    let el = this
    
  }
}
</script>

<style scoped>

</style>
