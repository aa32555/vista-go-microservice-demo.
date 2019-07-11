<template>
  <table class="table table-hover">
    <thead>
      <tr>
        <th v-for="key in columns" v-bind:key="key">
          {{ key }}
        </th>
      </tr>
    </thead>
    <tbody>
      <tr v-for="entry in rows" v-bind:key="entry.TableName" >
        <td v-for="key in columns" v-bind:key="key">
          {{entry[key]}}
        </td>
      </tr>
    </tbody>
  </table>
</template>

<script>
export default {
  name: 'HelloWorld',
  data () {
    return {
      'columns': ['Location', 'Name', 'Sex', 'Age', 'CheckOut', 'CheckIn'],
      'rows': [{'Location': '?', 'Name': '??', 'Sex': '?', 'Age': '?', 'CheckOut': '?', 'CheckIn': '?'}]
    }
  },
  created () {
    var _this = this
    $.getJSON('/api/', function (json) {
      _this.rows = json
    })
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1, h2 {
  font-weight: normal;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
