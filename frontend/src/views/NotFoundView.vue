<template>
  <div>
    <the-header />
    <base-card>
      <h1>404</h1>
      <p>The page you are looking for does not exist.</p>
    </base-card>
  </div>
</template>

<script lang="ts">
import { Vue } from 'vue-class-component';

export default class NotFoundView extends Vue {}
</script>
