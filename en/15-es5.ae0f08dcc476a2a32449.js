!function(){function e(){var t=c([":Select Language|Languages\u241fb47b79397919ae72f70bdc7252be4e238da67d8d\u241f1685538222344471199:",":START_TAG_ION_ITEM:English",":CLOSE_TAG_ION_ITEM:",":START_TAG_ION_ITEM_1:German",":CLOSE_TAG_ION_ITEM:"]);return e=function(){return t},t}function t(){var e=c([":Select Language|Header\u241f6729329861748a1ac918658f7871a97414091deb\u241f1461232356477253008:Select Language"]);return t=function(){return e},e}function n(){var e=c([":Settings|DrawerDocking\u241f54a38692b3cb01fd464eaff18a774cced79962ac\u241f4241466071733426227:Drawer Docking"]);return n=function(){return e},e}function o(){var e=c([":Settings|About\u241fe3ac5f8fae66ee7d38924f267b57dbc2ae4d16d8\u241f8383233070503389371:About"]);return o=function(){return e},e}function i(){var e=c([":Settings|Help\u241f38eca8ed965ed27304f3e8f5505492248159100b\u241f2306594643335405249:Help"]);return i=function(){return e},e}function r(){var e=c([":Settings|Navigate to Language Selection\u241f8943e53641a04eb7defe5a7e15698f4924c59196\u241f3167785507334604974:",":START_TAG_ION_LABEL:Language",":CLOSE_TAG_ION_LABEL:",":START_TAG_ION_NOTE:English",":CLOSE_TAG_ION_NOTE:"]);return r=function(){return e},e}function a(){var e=c([":Settings|Header\u241fcd3c4d3ddc94ff59b452b521a7945485f22a18cf\u241f2320123272127382813: Settings "]);return a=function(){return e},e}function c(e,t){return t||(t=e.slice(0)),Object.freeze(Object.defineProperties(e,{raw:{value:Object.freeze(t)}}))}function s(e,t){if(!(e instanceof t))throw new TypeError("Cannot call a class as a function")}function l(e,t){for(var n=0;n<t.length;n++){var o=t[n];o.enumerable=o.enumerable||!1,o.configurable=!0,"value"in o&&(o.writable=!0),Object.defineProperty(e,o.key,o)}}function u(e,t,n){return t&&l(e.prototype,t),n&&l(e,n),e}(window.webpackJsonp=window.webpackJsonp||[]).push([[15],{"7wo0":function(c,l,b){"use strict";b.r(l),b.d(l,"SettingsPageModule",function(){return y});var d=b("TEn/"),g=b("tyNb"),f=b("ofXK"),h=b("3Pt+"),p=b("mrSG"),m=b("AnxF"),v=b("KN1t"),L=b("YD5U"),k=b("fXoL"),O=b("e8h1"),w=["DrawerDockingToggle"];function T(e,t){if(1&e){var n=k.Mb();k.Lb(0,"ion-select",31),k.Xb("ionChange",function(){return k.jc(n),k.Zb().logoutFromMoodle()}),k.Lb(1,"ion-select-option",32),k.pc(2,"sign out"),k.Kb(),k.Kb()}if(2&e){var o=k.Zb();k.ec("interfaceOptions",o.actionSheetOptions)}}var S,K,E,_,M=((_=function(){function e(t,n,o){s(this,e),this.storage=t,this.moodleService=n,this.settingsModel=o,this.MOODLE_TOKEN="moodle_token",this.MOODLE_USER="moodle_user",this.actionSheetOptions={header:"Moodle"}}return u(e,[{key:"ngAfterViewInit",value:function(){this.drawerDockingToggle.checked=this.settingsModel.DrawerDocking}},{key:"ionViewWillEnter",value:function(){return Object(p.a)(this,void 0,void 0,regeneratorRuntime.mark(function e(){var t=this;return regeneratorRuntime.wrap(function(e){for(;;)switch(e.prev=e.next){case 0:this.storage.ready().then(function(){return Object(p.a)(t,void 0,void 0,regeneratorRuntime.mark(function e(){var t;return regeneratorRuntime.wrap(function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,this.isMoodleTokenPersisted();case 2:if(!e.sent){e.next=11;break}return e.next=5,this.moodleService.getCalenderEventsWeek(this.persistedMoodleToken).toPromise();case 5:if(t=e.sent,!this.moodleService.containsEvents(t)){e.next=11;break}return this.isSignedIn=!0,e.next=10,this.getMoodleUserName();case 10:return e.abrupt("return",void e.sent);case 11:this.setLoggedOutFromMoodle();case 12:case"end":return e.stop()}},e,this)}))});case 1:case"end":return e.stop()}},e,this)}))}},{key:"logoutFromMoodle",value:function(){return Object(p.a)(this,void 0,void 0,regeneratorRuntime.mark(function e(){var t=this;return regeneratorRuntime.wrap(function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,this.storage.remove(this.MOODLE_USER).then(function(){t.storage.remove(t.MOODLE_TOKEN).then(function(){t.setLoggedOutFromMoodle()})});case 2:case"end":return e.stop()}},e,this)}))}},{key:"isMoodleTokenPersisted",value:function(){return Object(p.a)(this,void 0,void 0,regeneratorRuntime.mark(function e(){var t=this;return regeneratorRuntime.wrap(function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,this.storage.get(this.MOODLE_TOKEN).then(function(e){return t.persistedMoodleToken=e,null!=t.persistedMoodleToken});case 2:return e.abrupt("return",e.sent);case 3:case"end":return e.stop()}},e,this)}))}},{key:"setLoggedOutFromMoodle",value:function(){this.isSignedIn=!1,this.moodleUserName="No user signed in."}},{key:"getMoodleUserName",value:function(){return Object(p.a)(this,void 0,void 0,regeneratorRuntime.mark(function e(){var t=this;return regeneratorRuntime.wrap(function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,this.storage.get(this.MOODLE_USER).then(function(e){t.moodleUserName=e});case 2:case"end":return e.stop()}},e,this)}))}},{key:"onDrawerDockingToggleChange",value:function(e){this.settingsModel.DrawerDocking=e.detail.checked,L.a.DrawerDocking=this.settingsModel.DrawerDocking}}]),e}()).\u0275fac=function(e){return new(e||_)(k.Ib(O.b),k.Ib(m.a),k.Ib(v.a))},_.\u0275cmp=k.Cb({type:_,selectors:[["app-settings"]],viewQuery:function(e,t){var n;1&e&&k.tc(w,!0),2&e&&k.gc(n=k.Yb())&&(t.drawerDockingToggle=n.first)},decls:48,vars:3,consts:function(){return[" Settings ",[1,"padding"],["lines","none"],[3,"routerLink"],["slot","start"],["size","large","name","person-circle-sharp"],[3,"innerHTML"],["selected-text","","interface","action-sheet","mode","ios",3,"interfaceOptions","ionChange",4,"ngIf"],["lines","",1,"ion-padding-top","ion-margin-top"],["routerLink","language"],"" + "\ufffd#21\ufffd" + "Language" + "\ufffd/#21\ufffd" + "" + "\ufffd#22\ufffd" + "English" + "\ufffd/#22\ufffd" + "",["slot","end"],["href","#"],["slot","end","color","dark","name","help-circle-outline"],"Help",["href","https://github.com/Studi-Guide","target","_blank"],["slot","end","color","dark","name","information-circle-outline"],"About",["slot","end","color","dark","name","information-circle-outline","data-tooltip","Einrasten der Schublade Ein-/Ausschalten","data-tooltip-pos","left center"],"Drawer Docking",[3,"ionChange"],["DrawerDockingToggle",""],[1,"ion-justify-content-center"],["size","large","name","logo-android","data-tooltip","Android","data-tooltip-pos","left center"],["size","large","name","logo-apple","data-tooltip","Apple","data-tooltip-pos","center top"],["size","large","name","logo-tux","data-tooltip","Linux","data-tooltip-pos","center bottom"],["size","large","name","logo-windows","data-tooltip","Windows"],["size","large","name","logo-firefox","data-tooltip","Firefox","data-tooltip-pos","center bottom"],["size","large","name","logo-chrome","data-tooltip","Chrome"],["size","large","name","logo-electron","data-tooltip","Electron","data-tooltip-pos","center bottom"],["size","large","name","logo-ionic","data-tooltip","Ionic","data-tooltip-pos","right center"],["selected-text","","interface","action-sheet","mode","ios",3,"interfaceOptions","ionChange"],["id","moodle-logout","value","sign-out","selected-text",""]]},template:function(e,t){1&e&&(k.Lb(0,"ion-header"),k.Lb(1,"ion-toolbar"),k.Lb(2,"ion-title"),k.Pb(3,0),k.Kb(),k.Kb(),k.Kb(),k.Lb(4,"ion-content"),k.Lb(5,"div",1),k.Lb(6,"ion-list-header"),k.pc(7,"Accounts"),k.Kb(),k.Lb(8,"ion-list",2),k.Lb(9,"ion-item",3),k.Lb(10,"ion-avatar",4),k.Jb(11,"ion-icon",5),k.Kb(),k.Lb(12,"ion-label"),k.Lb(13,"h3"),k.pc(14,"Moodle"),k.Kb(),k.Jb(15,"p",6),k.Kb(),k.oc(16,T,3,1,"ion-select",7),k.Kb(),k.Kb(),k.Kb(),k.Jb(17,"hr"),k.Lb(18,"ion-list",8),k.Lb(19,"ion-item",9),k.Sb(20,10),k.Jb(21,"ion-label"),k.Jb(22,"ion-note",11),k.Qb(),k.Kb(),k.Lb(23,"ion-item",12),k.Jb(24,"ion-icon",13),k.Lb(25,"ion-label"),k.Pb(26,14),k.Kb(),k.Kb(),k.Lb(27,"ion-item",15),k.Jb(28,"ion-icon",16),k.Lb(29,"ion-label"),k.Pb(30,17),k.Kb(),k.Kb(),k.Lb(31,"ion-item"),k.Jb(32,"ion-icon",18),k.Lb(33,"ion-label"),k.Pb(34,19),k.Kb(),k.Lb(35,"ion-toggle",20,21),k.Xb("ionChange",function(e){return t.onDrawerDockingToggleChange(e)}),k.Kb(),k.Kb(),k.Kb(),k.Kb(),k.Lb(37,"ion-footer"),k.Lb(38,"ion-grid"),k.Lb(39,"ion-row",22),k.Jb(40,"ion-icon",23),k.Jb(41,"ion-icon",24),k.Jb(42,"ion-icon",25),k.Jb(43,"ion-icon",26),k.Jb(44,"ion-icon",27),k.Jb(45,"ion-icon",28),k.Jb(46,"ion-icon",29),k.Jb(47,"ion-icon",30),k.Kb(),k.Kb(),k.Kb()),2&e&&(k.xb(9),k.fc("routerLink",t.isSignedIn?"/tabs/settings":"/tabs/schedule"),k.xb(6),k.ec("innerHTML",t.moodleUserName,k.kc),k.xb(1),k.ec("ngIf",t.isSignedIn))},directives:[d.v,d.V,d.T,d.p,d.C,d.B,d.y,d.cb,g.h,d.e,d.w,d.A,f.j,d.D,d.U,d.a,d.t,d.u,d.J,d.L,d.db,d.M],styles:[""]}),_),D=((E=function(){function e(t,n){s(this,e),this.router=t,this.locale=n,this.languages=[{Language:"English",Identifier:"en-US"},{Language:"German",Identifier:"de"}]}return u(e,[{key:"ngOnInit",value:function(){console.log(this.router.url,this.locale)}},{key:"SelectLanguageClick",value:function(e){console.log(e),e!==this.locale?this.router.navigate(["/"+e+this.router.url]):console.log("locale",this.locale,"not found in current URL")}}]),e}()).\u0275fac=function(e){return new(e||E)(k.Ib(g.g),k.Ib(k.v))},E.\u0275cmp=k.Cb({type:E,selectors:[["detail-view-select-language"]],decls:11,vars:0,consts:function(){var n,o;return n="Select Language",o="" + "\ufffd#9\ufffd" + "English" + "[\ufffd/#9\ufffd|\ufffd/#10\ufffd]" + "" + "\ufffd#10\ufffd" + "German" + "[\ufffd/#9\ufffd|\ufffd/#10\ufffd]" + "",[["translucent",""],["slot","start"],["defaultHref","/"],n,o=k.Rb(o),["href","/en"],["href","/de"]]},template:function(e,t){1&e&&(k.Lb(0,"ion-header",0),k.Lb(1,"ion-toolbar"),k.Lb(2,"ion-buttons",1),k.Jb(3,"ion-back-button",2),k.Kb(),k.Lb(4,"ion-title"),k.Pb(5,3),k.Kb(),k.Kb(),k.Kb(),k.Lb(6,"ion-content"),k.Lb(7,"ion-list"),k.Sb(8,4),k.Jb(9,"ion-item",5),k.Jb(10,"ion-item",6),k.Qb(),k.Kb(),k.Kb())},directives:[d.v,d.V,d.i,d.f,d.g,d.T,d.p,d.B,d.y],styles:[""]}),E),I=((K=function e(){s(this,e)}).\u0275mod=k.Gb({type:K}),K.\u0275inj=k.Fb({factory:function(e){return new(e||K)},imports:[[f.b,d.W]]}),K),y=((S=function e(){s(this,e)}).\u0275mod=k.Gb({type:S}),S.\u0275inj=k.Fb({factory:function(e){return new(e||S)},imports:[[d.W,f.b,h.a,g.i.forChild([{path:"",component:M},{path:"language",component:D}]),I]]}),S)}}])}();