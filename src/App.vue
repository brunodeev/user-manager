<template>
  <form>
    <input type="text" v-model="user.name">
    <input type="number" v-model="user.age">
    <button v-on:click="addUser">Adicionar</button>
  </form>

  <ul>
    <li v-for="user in users">{{ user.name }} - {{ user.age }}</li>
  </ul>
</template>

<script>
export default {
  data() {
    return {
      user: {
        name: '',
        age: '',
      },
      users: [],
    }
  },
  methods: {
    listUsers() {
      fetch('http://localhost:8080/users')
      .then(res => res.json())
      .then(body => {
        this.users = body
      })
    },
    addUser() {
      fetch('http://localhost:8080/users', {
        method: "POST",
        body: JSON.stringify(this.user),
        headers: {"Content-Type":"application/json; charset=UTF-8"},
      })
      .then(res => res.text())
      .then(body => {
        console.log("É -> " + body)
      });
    }
  },
  mounted(){
    this.listUsers();
  }
}
</script>