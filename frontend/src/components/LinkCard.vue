<template>
  <div class="tile is-parent" v-bind:class="{'is-hidden':hide}" >
  <div class="tile linkcard notification is-grey" style="cursor:pointer;" @click="redir">
    <div class="section">
    <p class="is-size-6">
    https://s.jlyon.org/{{link}}
    </p>
    <p class="is-size-7">
      {{val}}
    </p>
  </div>
</div>
  </div>
</template>

<script>
export default {
  props: ['val','text','link','img','myId','clicks','num'],
  name: 'LinkCard',
  data () {
    return{
      hide: true,
    }

  },
  methods:{
    redir: function(){
      location.href=this.link;
    }
  },
  mounted: function(){
    let el = this;
    setInterval(function(){
      el.hide = false;

    },
    50*el.num)
    // this.CardStyle.display = "inline-block";
  }
}
</script>

<style scoped>

.linkcard {
  -webkit-animation: animation 1000ms linear both;
  animation: animation 1000ms linear both;
}

/* Generated with Bounce.js. Edit at https://goo.gl/W7f9he */

@-webkit-keyframes animation {
  0% { -webkit-transform: matrix3d(0.5, 0, 0, 0, 0, 0.5, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1); transform: matrix3d(0.5, 0, 0, 0, 0, 0.5, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1); }
  4.3% { -webkit-transform: matrix3d(0.757, 0, 0, 0, 0, 0.757, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1); transform: matrix3d(0.757, 0, 0, 0, 0, 0.757, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1); }
  8.61% { -webkit-transform: matrix3d(0.939, 0, 0, 0, 0, 0.939, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1); transform: matrix3d(0.939, 0, 0, 0, 0, 0.939, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1); }
  12.91% { -webkit-transform: matrix3d(1.026, 0, 0, 0, 0, 1.026, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1); transform: matrix3d(1.026, 0, 0, 0, 0, 1.026, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1); }
  17.22% { -webkit-transform: matrix3d(1.047, 0, 0, 0, 0, 1.047, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1); transform: matrix3d(1.047, 0, 0, 0, 0, 1.047, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1); }
  28.33% { -webkit-transform: matrix3d(1.01, 0, 0, 0, 0, 1.01, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1); transform: matrix3d(1.01, 0, 0, 0, 0, 1.01, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1); }
  39.44% { -webkit-transform: matrix3d(0.997, 0, 0, 0, 0, 0.997, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1); transform: matrix3d(0.997, 0, 0, 0, 0, 0.997, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1); }
  61.66% { -webkit-transform: matrix3d(1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1); transform: matrix3d(1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1); }
  83.98% { -webkit-transform: matrix3d(1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1); transform: matrix3d(1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1); }
  100% { -webkit-transform: matrix3d(1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1); transform: matrix3d(1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1); }
}

@keyframes animation {
  0% { -webkit-transform: matrix3d(0.5, 0, 0, 0, 0, 0.5, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1); transform: matrix3d(0.5, 0, 0, 0, 0, 0.5, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1); }
  4.3% { -webkit-transform: matrix3d(0.757, 0, 0, 0, 0, 0.757, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1); transform: matrix3d(0.757, 0, 0, 0, 0, 0.757, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1); }
  8.61% { -webkit-transform: matrix3d(0.939, 0, 0, 0, 0, 0.939, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1); transform: matrix3d(0.939, 0, 0, 0, 0, 0.939, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1); }
  12.91% { -webkit-transform: matrix3d(1.026, 0, 0, 0, 0, 1.026, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1); transform: matrix3d(1.026, 0, 0, 0, 0, 1.026, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1); }
  17.22% { -webkit-transform: matrix3d(1.047, 0, 0, 0, 0, 1.047, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1); transform: matrix3d(1.047, 0, 0, 0, 0, 1.047, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1); }
  28.33% { -webkit-transform: matrix3d(1.01, 0, 0, 0, 0, 1.01, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1); transform: matrix3d(1.01, 0, 0, 0, 0, 1.01, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1); }
  39.44% { -webkit-transform: matrix3d(0.997, 0, 0, 0, 0, 0.997, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1); transform: matrix3d(0.997, 0, 0, 0, 0, 0.997, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1); }
  61.66% { -webkit-transform: matrix3d(1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1); transform: matrix3d(1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1); }
  83.98% { -webkit-transform: matrix3d(1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1); transform: matrix3d(1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1); }
  100% { -webkit-transform: matrix3d(1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1); transform: matrix3d(1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1); }
}



</style>
