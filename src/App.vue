<template>
  <form>
    <input type="text" v-model="user.name">
    <input type="number" v-model="user.age">
    <button v-on:click="addUser">Adicionar</button>
  </form>

  <ul>
    <li v-for="user in users">{{ user.name }}</li>
  </ul>
</template>

<script setup>
  import { ref } from 'vue';

  var users = ref([]);
  listUsers()

  var user = {
      name: '',
      age: '',
  }

  function listUsers() {
    fetch('http://localhost:8080/users')
    .then(res => res.json())
    .then(body => {
      users.value = body
    })
  }

  function addUser() {
    fetch('http://localhost:8080/users', {
      method: "POST",
      body: JSON.stringify(user),
      headers: {"Content-Type":"application/json; charset=UTF-8"},
    })
    .then(res => res.text())
    .then(body => {
      console.log("É -> " + body)
    });
  }
</script>