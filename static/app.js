var refresh = true;

var cre = Vue.component('create-button',{
  template: `
  <div style="width:100%;text-align:center;">
  <div v-if='show'><article-input></article-input></div>
  </div>`,
  data () {
    return{
      show: true,
    }
  },
  methods:{
    toggle: function(){
      this.show = !this.show;
    },
  }
})

var header = Vue.component('info-item',{
  props: ['val','text','link','img','myId','clicks'],
  template:`<div style="cursor: pointer; overflow: hide;" v-on:click="redir" class="box">
  <div class="media-content">
  <div>{{val.substr(0,35)}}...</div>
  <div>{{text}}</div>
  <div>Uses: {{clicks}} Link: /s/{{myId}}</div>
  </div>
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
  <div class="srx" style="margin: 0 auto; padding-bottom: 20px;">
  <input v-model="search" placeholder="search"></input>
  </div>
  <div class="container">
    <info-item v-for="article in filteredList" :clicks="article.count" :link="'/s/' + article.id" :img="article.image" :val="article.title" :text="article.text" :myId ="article.id"></info-item>
  </div>
  </div>`,
  data (){
    return {
      articles: [{title: "hey",text: "also hey",link: "asdf",img: ""}],
      search: "",
    }
  },
  computed: {
    filteredList() {
      return this.articles.filter(article => {
        return article.title.toLowerCase().includes(this.search.toLowerCase())
      })
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
  <input v-model="sendVal.link" placeholder="link"></input>
  <input v-model="sendVal.text" placeholder="🤖"></input>
  <button @click=submit>Add</button>
  <br><br>
  <a v-bind:href="link" style="font-size:20px;">{{data}}</a>
  </div>`,
  data (){
    return {
      sendVal: {title: "", text: "", image: "", link: "https://",password:""},
      data: "",
      link:""
    }
  },
  methods:{
    submit: function(){
      let el = this;
      this.sendVal.title = this.sendVal.link;
      $.post("/add",JSON.stringify(this.sendVal),function(data){
        el.data = "/s/";
        el.data += data;
        el.link = "/s/";
        el.link += data;

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


Vue.component("titlebar",{
  template: `<div v-bind:style=titleStyle><p v-bind:style=paragraphStyle>{{titleText}}</p></div>`,
  data (){
    return{
      titleStyle: {position:"absolute",backgroundColor:"#eee",height:"50px",width:"auto",top:"1",left:"0",right:"0"},
      paragraphStyle: {float: "left",height:"34px",lineHeight:"50px",verticalAlign:"center",paddingLeft:"30px",margin:"0"},
      titleText: "Shortlinks"
    }
  }

});

var App = new Vue({
  el: '#app-vue',
  data: {
  },


});
