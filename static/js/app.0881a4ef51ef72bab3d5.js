webpackJsonp([1],{"4NBF":function(e,t){},NHnr:function(e,t,n){"use strict";Object.defineProperty(t,"__esModule",{value:!0});var o=n("7+uW"),r={render:function(){var e=this.$createElement,t=this._self._c||e;return t("div",{attrs:{id:"app"}},[t("router-view")],1)},staticRenderFns:[]};var i=n("VU/8")({name:"App"},r,!1,function(e){n("Ospb")},null,null).exports,a=n("/ocq"),s=n("mvHQ"),l=n.n(s),c=n("13sD"),m=(n("fIPj"),n("6x4x")),u=n("AaoT"),d=(n("4NBF"),n("bQRx"));c.Terminal.applyAddon(d),c.Terminal.applyAddon(m),c.Terminal.applyAddon(u);var p=c.Terminal,f=n("xrTZ"),h={name:"Console",props:{terminal:{type:Object,default:function(){return[]}}},data:function(){return{term:null,terminalSocket:null}},methods:{runRealTerminal:function(){console.log("webSocket is finished")},closeRealTerminal:function(){console.log("close")}},mounted:function(){console.log("pid : "+this.terminal.pid+" is on ready");var e=document.getElementById("terminal");this.term=new p({rows:70,fontSize:16,cursorBlink:!0,cursorStyle:"bar"}),this.term._initialized=!0,this.term.attachCustomKeyEventHandler(function(e){86===e.keyCode&&e.ctrlKey&&this.terminalSocket.send((new TextEncoder).encode("\0"+this.copy))}),this.term.open(e),this.term.fit(),this.term.scrollToBottom(),this.terminalSocket=new WebSocket("ws://"+window.location.host+"/ssh/ws"),this.terminalSocket.onopen=this.runRealTerminal,this.terminalSocket.onclose=this.closeRealTerminal,this.terminalSocket.onerror=function(e){console.log("error",e)},function(e,t,n,o){e.socket=t;var r=null,i=function(t){o&&o>0?r?r+=t.data:(r=t.data,setTimeout(function(){e.write(r)},o)):e.write(t.data)},a=function(e){t.send(l()({type:"cmd",cmd:f.Base64.encode(e)}))};t.onmessage=i,n&&e.on("data",a);var s=setInterval(function(){t.send(l()({type:"heartbeat",data:""}))},2e4);t.addEventListener("close",function(){t.removeEventListener("message",i),e.off("data",a),delete e.socket,clearInterval(s)})}(this.term,this.terminalSocket,!0,-1),console.log("mounted is going on")},beforeDestroy:function(){this.terminalSocket.close(),this.term.disposal()}},v={render:function(){var e=this.$createElement;return(this._self._c||e)("div",{staticClass:"console",attrs:{id:"terminal"}})},staticRenderFns:[]},y={name:"WebSSH",data:function(){return{terminal:{pid:1,name:"terminal"}}},components:{"my-terminal":n("VU/8")(h,v,!1,null,null,null).exports}},S={render:function(){var e=this.$createElement,t=this._self._c||e;return t("div",{staticClass:"container"},[t("my-terminal",{attrs:{terminal:this.terminal}})],1)},staticRenderFns:[]};var k=n("VU/8")(y,S,!1,function(e){n("OEIC")},"data-v-5a6995eb",null).exports;o.a.use(a.a);var w=new a.a({routes:[{path:"/",name:"WebSSH",component:k}]});o.a.config.productionTip=!1,new o.a({el:"#app",router:w,components:{App:i},template:"<App/>"})},OEIC:function(e,t){},Ospb:function(e,t){},fIPj:function(e,t){}},["NHnr"]);
//# sourceMappingURL=app.0881a4ef51ef72bab3d5.js.map