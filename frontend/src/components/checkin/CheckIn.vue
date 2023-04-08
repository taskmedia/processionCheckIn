<template>
  <div class="hello">
    <h1>Not yet implemented: {{ name }}</h1>
    <base-card>
      <h2>Hello World</h2>
    </base-card>
  </div>
</template>

<script lang="ts">
import { Vue } from 'vue-class-component';

export default class CheckIn extends Vue {
  name = 'CheckIn';
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
</style>
