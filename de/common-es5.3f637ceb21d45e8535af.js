!function(){function n(n,t){for(var e=0;e<t.length;e++){var r=t[e];r.enumerable=r.enumerable||!1,r.configurable=!0,"value"in r&&(r.writable=!0),Object.defineProperty(n,r.key,r)}}function t(n,t,e,r,i,o,u){try{var a=n[o](u),c=a.value}catch(s){return void e(s)}a.done?t(c):Promise.resolve(c).then(r,i)}function e(n){return function(){var e=this,r=arguments;return new Promise(function(i,o){var u=n.apply(e,r);function a(n){t(u,i,o,a,c,"next",n)}function c(n){t(u,i,o,a,c,"throw",n)}a(void 0)})}}(window.webpackJsonp=window.webpackJsonp||[]).push([[0],{"0/6H":function(n,t,e){"use strict";e.d(t,"a",function(){return u});var r=e("A36C"),i=e("iWo5"),o=e("qULd"),u=function(n,t){var e,u,a=function(n,r,i){if("undefined"!=typeof document){var o=document.elementFromPoint(n,r);o&&t(o)?o!==e&&(s(),c(o,i)):s()}},c=function(n,t){e=n,u||(u=e);var i=e;Object(r.f)(function(){return i.classList.add("ion-activated")}),t()},s=function(){var n=arguments.length>0&&void 0!==arguments[0]&&arguments[0];if(e){var t=e;Object(r.f)(function(){return t.classList.remove("ion-activated")}),n&&u!==e&&e.click(),e=void 0}};return Object(i.createGesture)({el:n,gestureName:"buttonActiveDrag",threshold:0,onStart:function(n){return a(n.currentX,n.currentY,o.a)},onMove:function(n){return a(n.currentX,n.currentY,o.b)},onEnd:function(){s(!0),Object(o.e)(),u=void 0}})}},"74mu":function(n,t,r){"use strict";r.d(t,"a",function(){return o}),r.d(t,"b",function(){return u}),r.d(t,"c",function(){return i}),r.d(t,"d",function(){return c});var i=function(n,t){return null!==t.closest(n)},o=function(n,t){return"string"==typeof n&&n.length>0?Object.assign((i=!0,(r="ion-color-"+n)in(e={"ion-color":!0})?Object.defineProperty(e,r,{value:i,enumerable:!0,configurable:!0,writable:!0}):e[r]=i,e),t):t;var e,r,i},u=function(n){var t={};return function(n){return void 0!==n?(Array.isArray(n)?n:n.split(" ")).filter(function(n){return null!=n}).map(function(n){return n.trim()}).filter(function(n){return""!==n}):[]}(n).forEach(function(n){return t[n]=!0}),t},a=/^[a-z][a-z0-9+\-.]*:/,c=function(){var n=e(regeneratorRuntime.mark(function n(t,e,r,i){var o;return regeneratorRuntime.wrap(function(n){for(;;)switch(n.prev=n.next){case 0:if(null==t||"#"===t[0]||a.test(t)){n.next=4;break}if(!(o=document.querySelector("ion-router"))){n.next=4;break}return n.abrupt("return",(null!=e&&e.preventDefault(),o.push(t,r,i)));case 4:return n.abrupt("return",!1);case 5:case"end":return n.stop()}},n)}));return function(t,e,r,i){return n.apply(this,arguments)}}()},AnxF:function(t,e,r){"use strict";r.d(e,"a",function(){return a});var i=r("AytR"),o=r("fXoL"),u=r("tk/3"),a=function(){var t=function(){function t(n,e){!function(n,t){if(!(n instanceof t))throw new TypeError("Cannot call a class as a function")}(this,t),this.httpClient=n,this.env=e,this.moodleUrl="https://moodle3.de"}var e,r,i;return e=t,(r=[{key:"getLoginToken",value:function(n,t){var e=this.moodleUrl+"/login/token.php"+this.generateUrlQuery({username:n,password:t,service:"moodle_mobile_app"});return this.httpClient.get(e)}},{key:"getCalenderEventsMonth",value:function(n){return this.moodleRequest(n,"core_calendar_get_calendar_monthly_view",null)}},{key:"getCalenderEventsWeek",value:function(n){return this.moodleRequest(n,"core_calendar_get_calendar_upcoming_view",null)}},{key:"containsEvents",value:function(n){return null!=n.events||void 0!==n.events}},{key:"containsToken",value:function(n){return null!=n.token||void 0!==n.token}},{key:"moodleRequest",value:function(n,t,e){var r=this.moodleUrl+"/webservice/rest/server.php"+this.generateUrlQuery({wstoken:n.token,wsfunction:t,moodlewsrestformat:"json"});return this.httpClient.get(r)}},{key:"generateUrlQuery",value:function(n){if(null==n||0===Object.keys(n).length)return"";var t="?";for(var e in n)n.hasOwnProperty(e)&&("?"!==t&&(t+="&"),t+=e+"="+n[e]);return t}}])&&n(e.prototype,r),i&&n(e,i),t}();return t.\u0275fac=function(n){return new(n||t)(o.Tb(u.b),o.Tb(i.a))},t.\u0275prov=o.Eb({token:t,factory:t.\u0275fac,providedIn:"root"}),t}()},ZaV5:function(n,t,r){"use strict";r.d(t,"a",function(){return i}),r.d(t,"b",function(){return o});var i=function(){var n=e(regeneratorRuntime.mark(function n(t,e,r,i,o){var u;return regeneratorRuntime.wrap(function(n){for(;;)switch(n.prev=n.next){case 0:if(!t){n.next=2;break}return n.abrupt("return",t.attachViewToDom(e,r,o,i));case 2:if("string"==typeof r||r instanceof HTMLElement){n.next=4;break}throw new Error("framework delegate is missing");case 4:if(u="string"==typeof r?e.ownerDocument&&e.ownerDocument.createElement(r):r,i&&i.forEach(function(n){return u.classList.add(n)}),o&&Object.assign(u,o),e.appendChild(u),n.t0=u.componentOnReady,!n.t0){n.next=12;break}return n.next=12,u.componentOnReady();case 12:return n.abrupt("return",u);case 13:case"end":return n.stop()}},n)}));return function(t,e,r,i,o){return n.apply(this,arguments)}}(),o=function(n,t){if(t){if(n)return n.removeViewFromDom(t.parentElement,t);t.remove()}return Promise.resolve()}},h3R7:function(n,t,e){"use strict";e.d(t,"a",function(){return r});var r={bubbles:{dur:1e3,circles:9,fn:function(n,t,e){var r=n*t/e-n+"ms",i=2*Math.PI*t/e;return{r:5,style:{top:9*Math.sin(i)+"px",left:9*Math.cos(i)+"px","animation-delay":r}}}},circles:{dur:1e3,circles:8,fn:function(n,t,e){var r=t/e,i=n*r-n+"ms",o=2*Math.PI*r;return{r:5,style:{top:9*Math.sin(o)+"px",left:9*Math.cos(o)+"px","animation-delay":i}}}},circular:{dur:1400,elmDuration:!0,circles:1,fn:function(){return{r:20,cx:48,cy:48,fill:"none",viewBox:"24 24 48 48",transform:"translate(0,0)",style:{}}}},crescent:{dur:750,circles:1,fn:function(){return{r:26,style:{}}}},dots:{dur:750,circles:3,fn:function(n,t){return{r:6,style:{left:9-9*t+"px","animation-delay":-110*t+"ms"}}}},lines:{dur:1e3,lines:12,fn:function(n,t,e){return{y1:17,y2:29,style:{transform:"rotate(".concat(30*t+(t<6?180:-180),"deg)"),"animation-delay":n*t/e-n+"ms"}}}},"lines-small":{dur:1e3,lines:12,fn:function(n,t,e){return{y1:12,y2:20,style:{transform:"rotate(".concat(30*t+(t<6?180:-180),"deg)"),"animation-delay":n*t/e-n+"ms"}}}}}},qULd:function(n,t,e){"use strict";e.d(t,"a",function(){return o}),e.d(t,"b",function(){return u}),e.d(t,"c",function(){return i}),e.d(t,"d",function(){return c}),e.d(t,"e",function(){return a});var r={getEngine:function(){var n=window;return n.TapticEngine||n.Capacitor&&n.Capacitor.isPluginAvailable("Haptics")&&n.Capacitor.Plugins.Haptics},available:function(){return!!this.getEngine()},isCordova:function(){return!!window.TapticEngine},isCapacitor:function(){return!!window.Capacitor},impact:function(n){var t=this.getEngine();if(t){var e=this.isCapacitor()?n.style.toUpperCase():n.style;t.impact({style:e})}},notification:function(n){var t=this.getEngine();if(t){var e=this.isCapacitor()?n.style.toUpperCase():n.style;t.notification({style:e})}},selection:function(){this.impact({style:"light"})},selectionStart:function(){var n=this.getEngine();n&&(this.isCapacitor()?n.selectionStart():n.gestureSelectionStart())},selectionChanged:function(){var n=this.getEngine();n&&(this.isCapacitor()?n.selectionChanged():n.gestureSelectionChanged())},selectionEnd:function(){var n=this.getEngine();n&&(this.isCapacitor()?n.selectionEnd():n.gestureSelectionEnd())}},i=function(){r.selection()},o=function(){r.selectionStart()},u=function(){r.selectionChanged()},a=function(){r.selectionEnd()},c=function(n){r.impact(n)}}}])}();