<script setup lang="ts">
</script>

<template>
  <div class="view">
    <span class="heading">
      <slot name="heading" />
    </span>
    <slot />
    <div class="actions">
      <slot name="actions" />
    </div>
    <slot name="afterActions" />
  </div>
</template>

<style scoped lang="scss">
.view {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 22px;
  padding: 22px;
  color: rgba(var(--fg-rgb), .7);
  text-wrap: pretty;
  .heading {
    font-size: 32px;
    margin-bottom: 22px;
    color: var(--fg-color);
  }
  .actions {
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 22px;
  }
}
</style>