!function(){function e(e,t){var i;if("undefined"==typeof Symbol||null==e[Symbol.iterator]){if(Array.isArray(e)||(i=function(e,t){if(!e)return;if("string"==typeof e)return n(e,t);var i=Object.prototype.toString.call(e).slice(8,-1);"Object"===i&&e.constructor&&(i=e.constructor.name);if("Map"===i||"Set"===i)return Array.from(e);if("Arguments"===i||/^(?:Ui|I)nt(?:8|16|32)(?:Clamped)?Array$/.test(i))return n(e,t)}(e))||t&&e&&"number"==typeof e.length){i&&(e=i);var r=0,o=function(){};return{s:o,n:function(){return r>=e.length?{done:!0}:{done:!1,value:e[r++]}},e:function(e){throw e},f:o}}throw new TypeError("Invalid attempt to iterate non-iterable instance.\nIn order to be iterable, non-array objects must have a [Symbol.iterator]() method.")}var s,a=!0,c=!1;return{s:function(){i=e[Symbol.iterator]()},n:function(){var e=i.next();return a=e.done,e},e:function(e){c=!0,s=e},f:function(){try{a||null==i.return||i.return()}finally{if(c)throw s}}}}function n(e,n){(null==n||n>e.length)&&(n=e.length);for(var t=0,i=new Array(n);t<n;t++)i[t]=e[t];return i}function t(e,n){if(!(e instanceof n))throw new TypeError("Cannot call a class as a function")}function i(e,n){for(var t=0;t<n.length;t++){var i=n[t];i.enumerable=i.enumerable||!1,i.configurable=!0,"value"in i&&(i.writable=!0),Object.defineProperty(e,i.key,i)}}function r(e,n,t){return n&&i(e.prototype,n),t&&i(e,t),e}(window.webpackJsonp=window.webpackJsonp||[]).push([[15],{L0xO:function(n,i,o){"use strict";o.r(i),o.d(i,"SchedulePageModule",function(){return M});var s=o("TEn/"),a=o("tyNb"),c=o("ofXK"),u=o("3Pt+"),l=o("mrSG"),d=o("AnxF"),f=o("fXoL"),h=o("59pt"),b=o("KN1t"),g=o("e8h1");function v(e,n){if(1&e){var t=f.Mb();f.Lb(0,"div"),f.Lb(1,"ion-header"),f.Lb(2,"ion-toolbar"),f.Lb(3,"ion-title"),f.pc(4," Moodle Login "),f.Kb(),f.Kb(),f.Kb(),f.Lb(5,"ion-grid",1),f.Lb(6,"ion-row"),f.Lb(7,"ion-col",2),f.Lb(8,"ion-item"),f.Lb(9,"ion-label",3),f.pc(10,"user name"),f.Kb(),f.Lb(11,"ion-input",4),f.Xb("ionChange",function(){return f.jc(t),f.Zb().clearInvalidCredentialsMsg()})("ngModelChange",function(e){return f.jc(t),f.Zb().userInput=e}),f.Jb(12,"ion-icon",5),f.Kb(),f.Kb(),f.Lb(13,"ion-item"),f.Lb(14,"ion-label",3),f.pc(15,"password"),f.Kb(),f.Lb(16,"ion-input",6),f.Xb("ionChange",function(){return f.jc(t),f.Zb().clearInvalidCredentialsMsg()})("ngModelChange",function(e){return f.jc(t),f.Zb().userPassword=e}),f.Jb(17,"ion-icon",7),f.Kb(),f.Kb(),f.Lb(18,"ion-button",8),f.Xb("click",function(){return f.jc(t),f.Zb().fetchAndPersistMoodleToken()}),f.Lb(19,"h4"),f.pc(20,"sign in"),f.Kb(),f.Kb(),f.Kb(),f.Kb(),f.Lb(21,"ion-row"),f.Jb(22,"ion-col",9),f.Kb(),f.Kb(),f.Kb()}if(2&e){var i=f.Zb();f.xb(11),f.ec("ngModel",i.userInput),f.xb(5),f.ec("ngModel",i.userPassword),f.xb(6),f.ec("innerHTML",i.invalidCredentialsMessage,f.kc)}}var p,m=((p=function(){function e(n,i,r,o,s){t(this,e),this.storage=n,this.moodleService=i,this.faio=r,this.platform=o,this.settings=s,this.isSignedIn=new f.o,this.moodleToken=new f.o,this.isUserLoggedIn=!0,this.MOODLE_TOKEN="moodle_token",this.MOODLE_USER="moodle_user",this.userInput="admin",this.userPassword="administrator"}return r(e,[{key:"ngOnInit",value:function(){}},{key:"checkMoodleLoginState",value:function(){return Object(l.a)(this,void 0,void 0,regeneratorRuntime.mark(function e(){var n;return regeneratorRuntime.wrap(function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,this.storage.ready();case 2:if(n=!1,!(this.platform.is("hybrid")&&this.userInput.length>0&&this.userPassword.length>0)){e.next=14;break}return e.prev=4,e.next=7,this.faio.show({});case 7:n=!0,console.log("Face ID result"+n),e.next=14;break;case 11:e.prev=11,e.t0=e.catch(4),console.log(e.t0);case 14:if(e.t1=n,!e.t1){e.next=18;break}return e.next=18,this.fetchAndPersistMoodleToken();case 18:return e.next=20,this.getPersistedToken();case 20:this.isUserLoggedIn=null!=this.token,this.isSignedIn.emit(this.isUserLoggedIn),this.isUserLoggedIn&&this.moodleToken.emit(this.token);case 23:case"end":return e.stop()}},e,this,[[4,11]])}))}},{key:"fetchAndPersistMoodleToken",value:function(){return Object(l.a)(this,void 0,void 0,regeneratorRuntime.mark(function e(){var n,t,i;return regeneratorRuntime.wrap(function(e){for(;;)switch(e.prev=e.next){case 0:return n=this.userInput,t=this.userPassword,e.next=4,this.moodleService.getLoginToken(n,t).toPromise();case 4:if(i=e.sent,!this.moodleService.containsToken(i)){e.next=17;break}return this.isUserLoggedIn=!0,this.moodleToken.emit(i),this.isSignedIn.emit(this.isUserLoggedIn),e.next=11,this.storage.set(this.MOODLE_USER,n);case 11:return e.next=13,this.storage.set(this.MOODLE_TOKEN,i);case 13:this.userInput="",this.userPassword="",e.next=18;break;case 17:this.isUserLoggedIn=!1,this.isSignedIn.emit(this.isUserLoggedIn),this.invalidCredentialsMessage="Invalid credentials";case 18:case"end":return e.stop()}},e,this)}))}},{key:"clearInvalidCredentialsMsg",value:function(){this.invalidCredentialsMessage=""}},{key:"getPersistedToken",value:function(){return Object(l.a)(this,void 0,void 0,regeneratorRuntime.mark(function e(){var n=this;return regeneratorRuntime.wrap(function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,this.storage.get(this.MOODLE_TOKEN).then(function(e){n.token=e});case 2:case"end":return e.stop()}},e,this)}))}}]),e}()).\u0275fac=function(e){return new(e||p)(f.Ib(g.b),f.Ib(d.a),f.Ib(h.a),f.Ib(s.ab),f.Ib(b.a))},p.\u0275cmp=f.Cb({type:p,selectors:[["app-login"]],outputs:{isSignedIn:"isSignedIn",moodleToken:"moodleToken"},decls:1,vars:1,consts:[[4,"ngIf"],["fixed","md"],["size","12"],["position","stacked"],["type","text","clear-input","",1,"ion-text-center",3,"ngModel","ionChange","ngModelChange"],["name","person-outline"],["type","password","clear-input","",1,"ion-text-center",3,"ngModel","ionChange","ngModelChange"],["name","key-outline"],["size","default","expand","full","mode","ios","color","secondary",3,"click"],["size","12","id","invalid-credentials",1,"ion-text-center",3,"innerHTML"]],template:function(e,n){1&e&&f.oc(0,v,23,3,"div",0),2&e&&f.ec("ngIf",!n.isUserLoggedIn)},directives:[c.j,s.s,s.U,s.S,s.r,s.I,s.o,s.v,s.x,s.u,s.eb,u.d,u.e,s.t,s.h],styles:["ion-button[_ngcontent-%COMP%]{width:calc(100% - 22px);margin:12px}#invalid-credentials[_ngcontent-%COMP%]{color:#eb445a;font-size:20px}"]}),p);function k(e,n){if(1&e&&(f.Lb(0,"ion-card-title"),f.pc(1),f.Kb()),2&e){var t=f.Zb().$implicit;f.xb(1),f.qc(t.course.fullname)}}function x(e,n){if(1&e){var t=f.Mb();f.Lb(0,"ion-item-sliding"),f.Lb(1,"ion-item",6),f.Lb(2,"ion-card"),f.Lb(3,"ion-card-header"),f.oc(4,k,2,1,"ion-card-title",2),f.Jb(5,"ion-card-subtitle",7),f.Kb(),f.Lb(6,"ion-card-content"),f.Jb(7,"div",7),f.Lb(8,"ion-text",8),f.Xb("click",function(){f.jc(t);var e=n.$implicit;return f.Zb(2).onLocationClick(e.location)}),f.pc(9),f.Kb(),f.Kb(),f.Kb(),f.Kb(),f.Kb()}if(2&e){var i=n.$implicit;f.xb(4),f.ec("ngIf",i.course),f.xb(1),f.ec("innerHTML",i.name,f.kc),f.xb(2),f.ec("innerHTML",i.description,f.kc),f.xb(2),f.rc("Location: ",i.location,"")}}function y(e,n){if(1&e){var t=f.Mb();f.Lb(0,"div"),f.Lb(1,"ion-refresher",3),f.Xb("ionRefresh",function(e){return f.jc(t),f.Zb().doRefreshEvents(e)}),f.Jb(2,"ion-refresher-content"),f.Kb(),f.Lb(3,"ion-list",4),f.oc(4,x,10,4,"ion-item-sliding",5),f.Kb(),f.Kb()}if(2&e){var i=f.Zb();f.xb(4),f.ec("ngForOf",i.calenderEvents)}}var L,w,I=((w=function(){function n(e,i,r){t(this,n),this.moodleService=e,this.loadingController=i,this.router=r,this.calenderEvents=[]}return r(n,[{key:"ionViewWillEnter",value:function(){return Object(l.a)(this,void 0,void 0,regeneratorRuntime.mark(function e(){return regeneratorRuntime.wrap(function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,this.login.checkMoodleLoginState();case 2:case"end":return e.stop()}},e,this)}))}},{key:"onSignIn",value:function(e){return Object(l.a)(this,void 0,void 0,regeneratorRuntime.mark(function n(){return regeneratorRuntime.wrap(function(n){for(;;)switch(n.prev=n.next){case 0:this.isMoodleUserSignedIn=!!e;case 1:case"end":return n.stop()}},n,this)}))}},{key:"fetchMoodleData",value:function(n){return Object(l.a)(this,void 0,void 0,regeneratorRuntime.mark(function t(){var i,r,o,s;return regeneratorRuntime.wrap(function(t){for(;;)switch(t.prev=t.next){case 0:return t.next=2,this.loadingController.create({message:"Collecting moodle data..."});case 2:return i=t.sent,t.next=5,i.present();case 5:return this.token=n,t.next=8,this.moodleService.getCalenderEventsWeek(n).toPromise();case 8:if(r=t.sent,!this.moodleService.containsEvents(r)){t.next=17;break}this.calenderEvents=this.CleanupEvents(r.events),o=e(this.calenderEvents);try{for(o.s();!(s=o.n()).done;)s.value.location="KA.206"}catch(a){o.e(a)}finally{o.f()}return t.next=15,i.dismiss();case 15:t.next=21;break;case 17:return this.isMoodleUserSignedIn=!1,this.login.isUserLoggedIn=!1,t.next=21,i.dismiss();case 21:case"end":return t.stop()}},t,this)}))}},{key:"doRefreshEvents",value:function(n){return Object(l.a)(this,void 0,void 0,regeneratorRuntime.mark(function t(){var i,r,o;return regeneratorRuntime.wrap(function(t){for(;;)switch(t.prev=t.next){case 0:if(!this.isMoodleUserSignedIn){t.next=5;break}return t.next=3,this.moodleService.getCalenderEventsWeek(this.token).toPromise();case 3:if(i=t.sent,this.moodleService.containsEvents(i)){this.calenderEvents=i.events,r=e(this.calenderEvents);try{for(r.s();!(o=r.n()).done;)o.value.location="KA.206"}catch(s){r.e(s)}finally{r.f()}n.target.complete()}else this.isMoodleUserSignedIn=!1;case 5:case"end":return t.stop()}},t,this)}))}},{key:"onLocationClick",value:function(e){return Object(l.a)(this,void 0,void 0,regeneratorRuntime.mark(function n(){return regeneratorRuntime.wrap(function(n){for(;;)switch(n.prev=n.next){case 0:return n.next=2,this.router.navigate(["tabs/navigation/detail"],{queryParams:{location:e}});case 2:case"end":return n.stop()}},n,this)}))}},{key:"ngAfterViewInit",value:function(){}},{key:"CleanupEvents",value:function(n){var t,i=new RegExp("<img[^>]*?>","g"),r=e(n);try{for(r.s();!(t=r.n()).done;){var o=t.value;if(i.test(o.description)){var s,a=e(o.description.match(i));try{for(a.s();!(s=a.n()).done;){var c=s.value;o.description=o.description.replace(c,"")}}catch(u){a.e(u)}finally{a.f()}}}}catch(u){r.e(u)}finally{r.f()}return n}}]),n}()).\u0275fac=function(e){return new(e||w)(f.Ib(d.a),f.Ib(s.X),f.Ib(a.g))},w.\u0275cmp=f.Cb({type:w,selectors:[["app-schedule"]],viewQuery:function(e,n){var t;1&e&&f.tc(m,!0),2&e&&f.gc(t=f.Yb())&&(n.login=t.first)},decls:4,vars:1,consts:[["fixed",""],[3,"isSignedIn","moodleToken"],[4,"ngIf"],["slot","fixed",3,"ionRefresh"],["lines","none"],[4,"ngFor","ngForOf"],[1,"ion-no-padding"],[3,"innerHTML"],["id","location-entry",3,"click"]],template:function(e,n){1&e&&(f.Lb(0,"ion-content"),f.Lb(1,"ion-grid",0),f.Lb(2,"app-login",1),f.Xb("isSignedIn",function(e){return n.onSignIn(e)})("moodleToken",function(e){return n.fetchMoodleData(e)}),f.Kb(),f.oc(3,y,5,1,"div",2),f.Kb(),f.Kb()),2&e&&(f.xb(3),f.ec("ngIf",n.isMoodleUserSignedIn))},directives:[s.p,s.r,m,c.j,s.E,s.F,s.y,c.i,s.w,s.v,s.j,s.l,s.m,s.k,s.R,s.n],styles:[".welcome-card[_ngcontent-%COMP%]   img[_ngcontent-%COMP%]{max-height:35vh;overflow:hidden}#location-entry[_ngcontent-%COMP%]{color:#1e90ff;font-weight:700;cursor:pointer}"]}),w),M=((L=function e(){t(this,e)}).\u0275mod=f.Gb({type:L}),L.\u0275inj=f.Fb({factory:function(e){return new(e||L)},imports:[[s.V,c.b,u.a,a.i.forChild([{path:"",component:I},{path:"login/",component:I}])]]}),L)}}])}();