<template>
  <header :class="mode">
    <nav>
      <h1><router-link to="/">CheckIn</router-link></h1>
      <ul>
        <router-link to="/">TODO</router-link> |
        <router-link to="/admin" :v-if="isAdmin">Admin</router-link>
      </ul>
    </nav>
  </header>
</template>

<script lang="ts">
import { Options, Vue } from 'vue-class-component';

// add props
@Options({
  props: {
    mode: {
      type: String,
      default: '',
    },
  },
  methods: {
    isAdmin() {
      return this.$store.getters['isAdmin'];
    },
  },
})

export default class TheHeader extends Vue {
  mode!: string;
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
$primary-color: #ffe110;
$secondary-color: #000;

header {
  width: 100%;
  height: 5rem;
  background-color: $primary-color;
  display: flex;
  justify-content: center;
  align-items: center;
}

header a {
  text-decoration: none;
  color: $secondary-color;
  display: inline-block;
  padding: 0.75rem 1.5rem;
  border: 1px solid transparent;
}

a:active,
a:hover,
a.router-link-active {
  border: 1px solid $secondary-color;
}

h1 {
  margin: 0;
}

h1 a {
  color: $secondary-color;
  margin: 0;
}

h1 a:hover,
h1 a:active,
h1 a.router-link-active {
  border-color: transparent;
}

header nav {
  width: 90%;
  margin: auto;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

header ul {
  list-style: none;
  margin: 0;
  padding: 0;
  display: flex;
  justify-content: center;
  align-items: center;
}

li {
  margin: 0 0.5rem;
}

header.admin {
  background-color: $secondary-color;
}

header.admin a {
  color: $primary-color;
}

header.admin a:hover,
header.admin a:active,
header.admin a.router-link-active {
  border-color: $primary-color;
}
</style>
