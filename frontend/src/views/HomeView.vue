<template>
  <div>
    <the-header />
    <div class="home">
      <h1>Welcome to the HomeView!</h1>
      <CheckIn />
    </div>
  </div>
</template>

<script lang="ts">
import { Options, Vue } from 'vue-class-component';
import CheckIn from '../components/checkin/CheckIn.vue';

@Options({
  components: {
    CheckIn,
  },
})
export default class HomeView extends Vue {}
</script>
