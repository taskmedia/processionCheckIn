<template>
  <div>
    <the-header mode="admin" />
    <div class="home">
      <h1>Welcome to the AdminView</h1>
    </div>
  </div>
</template>

<script lang="ts">
import { Options, Vue } from 'vue-class-component';

@Options({
  components: {},
})
export default class AdminView extends Vue {}
</script>
