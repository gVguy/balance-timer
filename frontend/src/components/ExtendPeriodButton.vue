<script setup lang="ts">
import TooltipHint from './TooltipHint.vue'
import RoundPressAndHoldButton from './RoundPressAndHoldButton.vue'

defineProps<{
  title: string
  text: string
}>()

defineEmits(['confirm'])

</script>

<template>
  <TooltipHint v-bind="$props">
    <RoundPressAndHoldButton @confirm="$emit('confirm')">
      <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24">
        <path d="M13 7V9H15V17H17V7H13M11 13H9V15H7V13H5V11H7V9H9V11H11V13Z" />
      </svg>
    </RoundPressAndHoldButton>
  </TooltipHint>
</template>
