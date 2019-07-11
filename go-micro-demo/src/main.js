// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import 'bootstrap'
import 'bootstrap/dist/css/bootstrap.min.css'
import JQuery from 'jquery'

window.$ = JQuery

Vue.config.productionTip = false

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  components: { App },
  template: '<App/>',
  data: {
    columns: ['TableName', 'ColumnName', 'ColumnType'],
    rows: [{'TableName': 'Hello', 'ColumnName': 'World', 'ColumnType': 'varchar'}]
  }
})
