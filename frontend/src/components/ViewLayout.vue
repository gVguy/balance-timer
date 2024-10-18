<script setup lang="ts">
</script>

<template>
  <div class="view">
    <span class="heading">
      <slot name="heading" />
    </span>
    <div class="spacer" />
    <slot />
    <div class="spacer" />
    <slot name="after" />
  </div>
</template>

<style scoped lang="scss">
.view {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 22px;
  .heading {
    font-size: 32px;
  }
  .spacer {
  }
  .after {
    margin-top: 32px;
  }
}
</style>