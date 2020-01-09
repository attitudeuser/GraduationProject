// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import './lib/css'
import './lib/script'
import './lib/global'

import Vue from 'vue'
import App from './App'
import router from './router'
import EventBus from './lib/eventBus.js'
import axios from 'axios'
import Vant from 'vant'
import ElementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css';
import "vant/lib/index.css";
Vue.use(ElementUI);
Vue.use(Vant);
Vue.prototype.$bus = EventBus
Vue.prototype.$http = axios
Vue.prototype.$api = 'http://localhost:7125'

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  template: '<App/>',
  components: {
    App
  }
})
