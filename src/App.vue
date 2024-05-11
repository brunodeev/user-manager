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

<template>
  <div class="container">
    <form class="forms">
      <div class="inputs">
        <label for="input-name">Nome</label>
        <input class="input-container" id="input-name" type="text" v-model="user.name">
      </div>
      <div class="inputs">
        <label for="input-age">Idade</label>
        <input class="input-container" id="input-age" type="number" v-model="user.age">
      </div>
      <button class="input-button" v-on:click="addUser">Adicionar</button>
    </form>

    <ul class="list">
      <li class="item-list" v-for="user in users">{{ user.name }} - {{ user.age }}</li>
    </ul>
  </div>
</template>

<style scoped>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
  font-family: 'Ubuntu Sans', sans-serif;
}

.container {
  background-color: #2e2e2e;
  height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
}

.forms {
  background-color: #f6f6f6;
  display: flex;
  flex-direction: column;
  justify-content: space-evenly;
  height: 20rem;
  width: 25rem;
  padding: 20px;
  border-radius: 12px;
  box-shadow: 0 0 20px rgba(0, 0, 0, .5);
}

.inputs {
  display: flex;
  flex-direction: column
}

.inputs label {
  line-height: 1.8;
}

.input-container {
  height: 2.5rem;
  border: none;
  border-radius: 8px;
  padding-left: 10px;
  box-shadow: 0 0 10px rgba(0, 0, 0, .1);
}

.input-button {
  height: 2.5rem;
  border: none;
  border-radius: 8px;
  background-color: #28cb61;
  cursor: pointer;
}
</style>