!function(){function e(){var t=d([":Appearance|Dark\u241f947b393e1e409c930df4857cf91075af77c5b354\u241f7232184396737982364:Dark"]);return e=function(){return t},t}function t(){var e=d([":Appearance|Light\u241f3cc73d43688ba9cd210df5ad4c795bfb29799b07\u241f274095433690503049:Light"]);return t=function(){return e},e}function n(){var e=d([":Appearance|System\u241fbf966718c45207ffb29f0530b27567c2f3fedf71\u241f8730898933626838300:System"]);return n=function(){return e},e}function o(){var e=d([":Appearance|Select Appearance\u241f1edc8b24c18f495f13f116bd42b3ade124b93f48\u241f3259427562907869969:Select Appearance"]);return o=function(){return e},e}function i(){var e=d([":Select Language|Languages\u241fb47b79397919ae72f70bdc7252be4e238da67d8d\u241f1685538222344471199:",":START_TAG_ION_ITEM:English",":CLOSE_TAG_ION_ITEM:",":START_TAG_ION_ITEM_1:German",":CLOSE_TAG_ION_ITEM:"]);return i=function(){return e},e}function r(){var e=d([":Select Language|Header\u241f6729329861748a1ac918658f7871a97414091deb\u241f1461232356477253008:Select Language"]);return r=function(){return e},e}function a(){var e=d([":Settings|About\u241fe3ac5f8fae66ee7d38924f267b57dbc2ae4d16d8\u241f8383233070503389371:About"]);return a=function(){return e},e}function c(){var e=d([":Settings|Help\u241f38eca8ed965ed27304f3e8f5505492248159100b\u241f2306594643335405249:Help"]);return c=function(){return e},e}function s(){var e=d([":Settings|DrawerDocking\u241f54a38692b3cb01fd464eaff18a774cced79962ac\u241f4241466071733426227:Drawer Docking"]);return s=function(){return e},e}function l(){var e=d([":Settings|Appearance\u241f6ec655cefb257488df55ae0cdc01eacb0157f4cc\u241f3826230940028795895:Appearance"]);return l=function(){return e},e}function u(){var e=d([":Settings|Navigate to Language Selection\u241fa4ca0895b1eb6a25c7e7f4fbc3808fd05e15ab4e\u241f4316220740899106338:",":START_TAG_ION_LABEL:Language",":CLOSE_TAG_ION_LABEL:",":START_TAG_ION_NOTE:English",":CLOSE_TAG_ION_NOTE:"]);return u=function(){return e},e}function b(){var e=d([":Settings|Header\u241fcd3c4d3ddc94ff59b452b521a7945485f22a18cf\u241f2320123272127382813: Settings "]);return b=function(){return e},e}function d(e,t){return t||(t=e.slice(0)),Object.freeze(Object.defineProperties(e,{raw:{value:Object.freeze(t)}}))}function f(e,t){if(!(e instanceof t))throw new TypeError("Cannot call a class as a function")}function g(e,t){for(var n=0;n<t.length;n++){var o=t[n];o.enumerable=o.enumerable||!1,o.configurable=!0,"value"in o&&(o.writable=!0),Object.defineProperty(e,o.key,o)}}function p(e,t,n){return t&&g(e.prototype,t),n&&g(e,n),e}(window.webpackJsonp=window.webpackJsonp||[]).push([[15],{"7wo0":function(d,g,h){"use strict";h.r(g),h.d(g,"SettingsPageModule",function(){return N});var m=h("TEn/"),v=h("tyNb"),L=h("ofXK"),k=h("3Pt+"),K=h("mrSG"),M=h("AnxF"),O=h("KN1t"),w=h("fXoL"),A=h("e8h1"),S=["DrawerDockingToggle"];function y(e,t){if(1&e){var n=w.Mb();w.Lb(0,"ion-select",35),w.Xb("ionChange",function(){return w.jc(n),w.Zb().logoutFromMoodle()}),w.Lb(1,"ion-select-option",36),w.pc(2,"sign out"),w.Kb(),w.Kb()}if(2&e){var o=w.Zb();w.ec("interfaceOptions",o.actionSheetOptions)}}var T,D,E,_,x,I=((x=function(){function e(t,n,o){f(this,e),this.storage=t,this.moodleService=n,this.settingsModel=o,this.MOODLE_TOKEN="moodle_token",this.MOODLE_USER="moodle_user",this.actionSheetOptions={header:"Moodle"}}return p(e,[{key:"ngAfterViewInit",value:function(){this.drawerDockingToggle.checked=this.settingsModel.DrawerDocking}},{key:"ionViewWillEnter",value:function(){return Object(K.a)(this,void 0,void 0,regeneratorRuntime.mark(function e(){var t=this;return regeneratorRuntime.wrap(function(e){for(;;)switch(e.prev=e.next){case 0:this.storage.ready().then(function(){return Object(K.a)(t,void 0,void 0,regeneratorRuntime.mark(function e(){var t;return regeneratorRuntime.wrap(function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,this.isMoodleTokenPersisted();case 2:if(!e.sent){e.next=11;break}return e.next=5,this.moodleService.getCalenderEventsWeek(this.persistedMoodleToken).toPromise();case 5:if(t=e.sent,!this.moodleService.containsEvents(t)){e.next=11;break}return this.isSignedIn=!0,e.next=10,this.getMoodleUserName();case 10:return e.abrupt("return",void e.sent);case 11:this.setLoggedOutFromMoodle();case 12:case"end":return e.stop()}},e,this)}))});case 1:case"end":return e.stop()}},e,this)}))}},{key:"logoutFromMoodle",value:function(){return Object(K.a)(this,void 0,void 0,regeneratorRuntime.mark(function e(){var t=this;return regeneratorRuntime.wrap(function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,this.storage.remove(this.MOODLE_USER).then(function(){t.storage.remove(t.MOODLE_TOKEN).then(function(){t.setLoggedOutFromMoodle()})});case 2:case"end":return e.stop()}},e,this)}))}},{key:"isMoodleTokenPersisted",value:function(){return Object(K.a)(this,void 0,void 0,regeneratorRuntime.mark(function e(){var t=this;return regeneratorRuntime.wrap(function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,this.storage.get(this.MOODLE_TOKEN).then(function(e){return t.persistedMoodleToken=e,null!=t.persistedMoodleToken});case 2:return e.abrupt("return",e.sent);case 3:case"end":return e.stop()}},e,this)}))}},{key:"setLoggedOutFromMoodle",value:function(){this.isSignedIn=!1,this.moodleUserName="No user signed in."}},{key:"getMoodleUserName",value:function(){return Object(K.a)(this,void 0,void 0,regeneratorRuntime.mark(function e(){var t=this;return regeneratorRuntime.wrap(function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,this.storage.get(this.MOODLE_USER).then(function(e){t.moodleUserName=e});case 2:case"end":return e.stop()}},e,this)}))}},{key:"onDrawerDockingToggleChange",value:function(e){this.settingsModel.DrawerDocking=e.detail.checked}},{key:"Appearance",get:function(){return this.settingsModel.AutoDarkMode?"System":this.settingsModel.DarkMode?"Dark":"Light"}}]),e}()).\u0275fac=function(e){return new(e||x)(w.Ib(A.b),w.Ib(M.a),w.Ib(O.a))},x.\u0275cmp=w.Cb({type:x,selectors:[["app-settings"]],viewQuery:function(e,t){var n;1&e&&w.tc(S,!0),2&e&&w.gc(n=w.Yb())&&(t.drawerDockingToggle=n.first)},decls:53,vars:4,consts:function(){return[" Settings ",["fixed",""],[1,"padding"],["lines","none"],[3,"routerLink"],["slot","start"],["size","large","name","person-circle-sharp"],[3,"innerHTML"],["selected-text","","interface","action-sheet","mode","ios",3,"interfaceOptions","ionChange",4,"ngIf"],["lines","",1,"ion-padding-top","ion-margin-top"],["routerLink","language"],"" + "\ufffd#22\ufffd" + "Language" + "\ufffd/#22\ufffd" + "" + "\ufffd#23\ufffd" + "English" + "\ufffd/#23\ufffd" + "",["slot","end"],["routerLink","appearance"],"Appearance",["slot","end",3,"textContent"],["slot","end","color","dark","name","information-circle-outline","data-tooltip","Einrasten der Schublade Ein-/Ausschalten","data-tooltip-pos","left center"],"Drawer Docking",[3,"ionChange"],["DrawerDockingToggle",""],["href","#"],["slot","end","color","dark","name","help-circle-outline"],"Help",["href","https://github.com/Studi-Guide","target","_blank"],["slot","end","color","dark","name","information-circle-outline"],"About",[1,"ion-justify-content-center"],["size","large","name","logo-android","data-tooltip","Android","data-tooltip-pos","left center"],["size","large","name","logo-apple","data-tooltip","Apple","data-tooltip-pos","center top"],["size","large","name","logo-tux","data-tooltip","Linux","data-tooltip-pos","center bottom"],["size","large","name","logo-windows","data-tooltip","Windows"],["size","large","name","logo-firefox","data-tooltip","Firefox","data-tooltip-pos","center bottom"],["size","large","name","logo-chrome","data-tooltip","Chrome"],["size","large","name","logo-electron","data-tooltip","Electron","data-tooltip-pos","center bottom"],["size","large","name","logo-ionic","data-tooltip","Ionic","data-tooltip-pos","right center"],["selected-text","","interface","action-sheet","mode","ios",3,"interfaceOptions","ionChange"],["id","moodle-logout","value","sign-out","selected-text",""]]},template:function(e,t){1&e&&(w.Lb(0,"ion-header"),w.Lb(1,"ion-toolbar"),w.Lb(2,"ion-title"),w.Pb(3,0),w.Kb(),w.Kb(),w.Kb(),w.Lb(4,"ion-content"),w.Lb(5,"ion-grid",1),w.Lb(6,"div",2),w.Lb(7,"ion-list-header"),w.pc(8,"Accounts"),w.Kb(),w.Lb(9,"ion-list",3),w.Lb(10,"ion-item",4),w.Lb(11,"ion-avatar",5),w.Jb(12,"ion-icon",6),w.Kb(),w.Lb(13,"ion-label"),w.Lb(14,"h3"),w.pc(15,"Moodle"),w.Kb(),w.Jb(16,"p",7),w.Kb(),w.oc(17,y,3,1,"ion-select",8),w.Kb(),w.Kb(),w.Kb(),w.Jb(18,"hr"),w.Lb(19,"ion-list",9),w.Lb(20,"ion-item",10),w.Sb(21,11),w.Jb(22,"ion-label"),w.Jb(23,"ion-note",12),w.Qb(),w.Kb(),w.Lb(24,"ion-item",13),w.Lb(25,"ion-label"),w.Pb(26,14),w.Kb(),w.Jb(27,"ion-note",15),w.Kb(),w.Lb(28,"ion-item"),w.Jb(29,"ion-icon",16),w.Lb(30,"ion-label"),w.Pb(31,17),w.Kb(),w.Lb(32,"ion-toggle",18,19),w.Xb("ionChange",function(e){return t.onDrawerDockingToggleChange(e)}),w.Kb(),w.Kb(),w.Lb(34,"ion-item",20),w.Jb(35,"ion-icon",21),w.Lb(36,"ion-label"),w.Pb(37,22),w.Kb(),w.Kb(),w.Lb(38,"ion-item",23),w.Jb(39,"ion-icon",24),w.Lb(40,"ion-label"),w.Pb(41,25),w.Kb(),w.Kb(),w.Kb(),w.Kb(),w.Kb(),w.Lb(42,"ion-footer"),w.Lb(43,"ion-grid"),w.Lb(44,"ion-row",26),w.Jb(45,"ion-icon",27),w.Jb(46,"ion-icon",28),w.Jb(47,"ion-icon",29),w.Jb(48,"ion-icon",30),w.Jb(49,"ion-icon",31),w.Jb(50,"ion-icon",32),w.Jb(51,"ion-icon",33),w.Jb(52,"ion-icon",34),w.Kb(),w.Kb(),w.Kb()),2&e&&(w.xb(10),w.fc("routerLink",t.isSignedIn?"/tabs/settings":"/tabs/schedule"),w.xb(6),w.ec("innerHTML",t.moodleUserName,w.kc),w.xb(1),w.ec("ngIf",t.isSignedIn),w.xb(10),w.ec("textContent",t.Appearance))},directives:[m.v,m.X,m.V,m.p,m.u,m.C,m.B,m.y,m.fb,v.h,m.e,m.w,m.A,L.j,m.D,m.W,m.a,m.t,m.L,m.N,m.gb,m.O],styles:[""]}),x),C=((_=function(){function e(t,n){f(this,e),this.router=t,this.locale=n,this.languages=[{Language:"English",Identifier:"en-US"},{Language:"German",Identifier:"de"}]}return p(e,[{key:"ngOnInit",value:function(){console.log(this.router.url,this.locale)}},{key:"SelectLanguageClick",value:function(e){console.log(e),e!==this.locale?this.router.navigate(["/"+e+this.router.url]):console.log("locale",this.locale,"not found in current URL")}}]),e}()).\u0275fac=function(e){return new(e||_)(w.Ib(v.g),w.Ib(w.v))},_.\u0275cmp=w.Cb({type:_,selectors:[["detail-view-select-language"]],decls:11,vars:0,consts:function(){var e,t;return e="Select Language",t="" + "\ufffd#9\ufffd" + "English" + "[\ufffd/#9\ufffd|\ufffd/#10\ufffd]" + "" + "\ufffd#10\ufffd" + "German" + "[\ufffd/#9\ufffd|\ufffd/#10\ufffd]" + "",[["translucent",""],["slot","start"],["defaultHref","/"],e,t=w.Rb(t),["href","/en"],["href","/de"]]},template:function(e,t){1&e&&(w.Lb(0,"ion-header",0),w.Lb(1,"ion-toolbar"),w.Lb(2,"ion-buttons",1),w.Jb(3,"ion-back-button",2),w.Kb(),w.Lb(4,"ion-title"),w.Pb(5,3),w.Kb(),w.Kb(),w.Kb(),w.Lb(6,"ion-content"),w.Lb(7,"ion-list"),w.Sb(8,4),w.Jb(9,"ion-item",5),w.Jb(10,"ion-item",6),w.Qb(),w.Kb(),w.Kb())},directives:[m.v,m.X,m.i,m.f,m.g,m.V,m.p,m.B,m.y],styles:[""]}),_),J=((E=function e(){f(this,e)}).\u0275mod=w.Gb({type:E}),E.\u0275inj=w.Fb({factory:function(e){return new(e||E)},imports:[[L.b,m.Y]]}),E),z=((D=function(){function e(t){f(this,e),this.settingsModel=t,this.selectedAppearance=t.AutoDarkMode?"system":t.DarkMode?"dark":"light"}return p(e,[{key:"ngOnInit",value:function(){}},{key:"onAppearanceChanged",value:function(e){var t;if(null===(t=null==e?void 0:e.detail)||void 0===t?void 0:t.value)switch(e.detail.value){case"system":this.settingsModel.AutoDarkMode=!0;break;case"light":this.settingsModel.AutoDarkMode&&(this.settingsModel.AutoDarkMode=!1),this.settingsModel.DarkMode=!1;break;case"dark":this.settingsModel.AutoDarkMode&&(this.settingsModel.AutoDarkMode=!1),this.settingsModel.DarkMode=!0}}}]),e}()).\u0275fac=function(e){return new(e||D)(w.Ib(O.a))},D.\u0275cmp=w.Cb({type:D,selectors:[["app-detail-view-select-appearance"]],decls:22,vars:1,consts:function(){return[["translucent",""],["slot","start"],["defaultHref","/"],"Select Appearance",["fixed",""],[3,"value","ionChange"],"System",["slot","end","color","primary","value","system"],"Light",["slot","end","color","primary","value","light"],"Dark",["slot","end","color","primary","value","dark"]]},template:function(e,t){1&e&&(w.Lb(0,"ion-header",0),w.Lb(1,"ion-toolbar"),w.Lb(2,"ion-buttons",1),w.Jb(3,"ion-back-button",2),w.Kb(),w.Lb(4,"ion-title"),w.Pb(5,3),w.Kb(),w.Kb(),w.Kb(),w.Lb(6,"ion-content"),w.Lb(7,"ion-grid",4),w.Lb(8,"ion-list"),w.Lb(9,"ion-radio-group",5),w.Xb("ionChange",function(e){return t.onAppearanceChanged(e)}),w.Lb(10,"ion-item"),w.Lb(11,"ion-label"),w.Pb(12,6),w.Kb(),w.Jb(13,"ion-radio",7),w.Kb(),w.Lb(14,"ion-item"),w.Lb(15,"ion-label"),w.Pb(16,8),w.Kb(),w.Jb(17,"ion-radio",9),w.Kb(),w.Lb(18,"ion-item"),w.Lb(19,"ion-label"),w.Pb(20,10),w.Kb(),w.Jb(21,"ion-radio",11),w.Kb(),w.Kb(),w.Kb(),w.Kb(),w.Kb()),2&e&&(w.xb(9),w.ec("value",t.selectedAppearance))},directives:[m.v,m.X,m.i,m.f,m.g,m.V,m.p,m.u,m.B,m.G,m.gb,m.y,m.A,m.F,m.eb],styles:[""]}),D),N=((T=function e(){f(this,e)}).\u0275mod=w.Gb({type:T}),T.\u0275inj=w.Fb({factory:function(e){return new(e||T)},imports:[[m.Y,L.b,k.a,v.i.forChild([{path:"",component:I},{path:"language",component:C},{path:"appearance",component:z}]),J]]}),T)}}])}();