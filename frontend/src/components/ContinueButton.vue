<script lang="ts" setup>

</script>

<template>
  <button class="continue-button">
    <svg class="triangle" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24">
      <path d="M2.5,1V23L21,12Z" />
    </svg>
  </button>
</template>

<style scoped lang="scss">
.continue-button {
  border-radius: 50%;
  width: 122px;
  height: 122px;
  overflow: hidden;
  .triangle {
    fill: var(--bg-color);
    width: 24px;
    height: 24px;
    pointer-events: none;
    position: relative;
    left: 3px;
  }
  &:hover {
    .triangle {
      animation: triangle-fly .4s;
      @keyframes triangle-fly {
        50% { translate: 72px; }
        51% { translate: -72px; }
      }
    }
  }
}
</style>
