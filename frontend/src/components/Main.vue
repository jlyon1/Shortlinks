<template>
  <div>
  <div class="section" style="backgroundColor:#111">
    <div class="subtitle has-text-white">
      Shortlinks | <a href="https://jlyon.org" style="color:white; textDecoration:none;">jlyon.org</a>
    </div>
  </div>
  <div class = "section">
    <div class="container">
  <div v-for="i in Math.ceil(links.length/3)" class="tile is-ancestor">
    <LinkCard v-for="(link,j) in links.slice((i - 1) * 3, i * 3)" :num=i+j :clicks="link.count" :link="link.id" :img="link.image" :val="link.title" :text="link.text" :myId ="link.id"></LinkCard>
  </div>
  </div>
  </div>
  </div>
</template>

<script>
import LinkCard from '@/components/LinkCard'


export default {
  name: 'Main',
  components: {LinkCard},
  data () {
    return {
      links: []
    }
  },
  mounted: function(){
    let el = this
    fetch("get/").then(res => {
     return (res.json())
    }).then(ret =>{
    el.links = ret;
    })
  }
}
</script>

<style scoped>

</style>
