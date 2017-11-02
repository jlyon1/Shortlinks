var refresh = true;

var header = Vue.component('info-item',{
  props: ['val','text','link','img','myId'],
  template:`<div style="display:none;" v-on:click="redir" class="info_item">
  <img v-if="img != ''" :src="img"></img>
  <span class="projectHeader">{{val}}</span>
  <div style="color:gray; float:right; font-size:12px;padding-top:5px; padding-right:5px;">{{myId}}</div>
  <span class="text">{{text}}</span>
  </div>`,
  data (){
    return {
    }
  },
  methods:{
    redir: function(){
      location.href=this.link;
    }
  },
  mounted(){
    $(".info_item").fadeIn("slow");
  }

})
var bdy = Vue.component('item-area',{
  template:`<div>
  <info-item v-for="article in articles" :link="article.link" :img="article.image" :val="article.title" :text="article.text" :myId ="article.id"></info-item>
  </div>`,
  data (){
    return {
      articles: [{title: "hey",text: "also hey",link: "asdf",img: ""}]
    }
  },
  mounted (){
    let el = this
    $.get("get/",function(data){
      el.articles = data
    })
    setInterval(function(){
      if(refresh){
        $.get("get/",function(data){
          el.articles = data
        })
        refresh = false;
      }
    },300)

  }

})

var articleIn = Vue.component('article-input',{
  props: ['msg'],
  template:`<div class ="ai">
  <input v-model="sendVal.title" placeholder="title"></input>
  <input v-model="sendVal.link" placeholder="link"></input>
  <input v-model="sendVal.password" placeholder="password" type="password"></input>
  <button @click=submit>Add</button>

  </div>`,
  data (){
    return {
      sendVal: {title: "", text: "", image: "", link: "",password:""}
    }
  },
  methods:{
    submit: function(){
      console.log("asdf");
      $.post("/add",JSON.stringify(this.sendVal),function(data){

      })
      refresh = true;
      this.sendVal.title = "";
      this.sendVal.text = "";
      this.sendVal.image ="";
      this.sendVal.link = "";
      this.sendVal.password = "";

    }
  }

})

var header = Vue.component('header-area',{
  props: ['msg'],
  template:`<div class="header">{{msg}}</div>`,
  data (){
    return {
    }
  }

})


var titlebar = Vue.component('title-bar',{
  props: ['msg'],
  template:`<div class="titlebar" style="cursor:pointer;" v-on:click="redir">{{msg}}</div>`,
  data (){
    return {
    }
  },
  methods: {
    redir: function(){
      location.href="https://jlyon.org";
    }
  }

})

var App = new Vue({
  el: '#app-vue',
  data: {
  }

});
