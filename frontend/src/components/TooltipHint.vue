<script setup lang="ts">
defineProps<{
  text?: string
  title?: string
}>()
</script>

<template>
  <div class="tooltip-wrapper">
    <slot />
    <div v-if="text || title" class="tooltip">
      <span v-if="title" class="title">{{ title }}</span>
      <span v-if="text" class="text">{{ text }}</span>
    </div>
  </div>
</template>

<style scoped lang="scss">
.tooltip-wrapper {
  position: relative;
  &:hover:not(:active) .tooltip {
    opacity: 1;
    transition-delay: .5s;
  }
  .tooltip {
    opacity: 0;
    transition: opacity .4s;
    pointer-events: none;
    --tooltip-border: 1px rgba(var(--fg-rgb), .3) solid;
    width: 120px;
    background: var(--bg-color);
    border: var(--tooltip-border);
    color: rgb(var(--fg-rgb), .7);
    text-align: center;
    padding: 8px;
    border-radius: 6px;
    width: 222px;
    position: absolute;
    top: -12px;
    left: 50%;
    translate: -50% -100%;
    display: flex;
    flex-direction: column;
    gap: 4px;
    &::after {
      content: '';
      position: absolute;
      height: 8px;
      width: 8px;
      top: 100%;
      left: 50%;
      translate: -50% -50%;
      rotate: 45deg;
      border-right: var(--tooltip-border);
      border-bottom: var(--tooltip-border);
      background: var(--bg-color);
    }
    .title {
      font-weight: bold;
    }
  }
}
</style>