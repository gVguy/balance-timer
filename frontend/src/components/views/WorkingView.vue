<script lang="ts" setup>
import CircleProgress from '@/components/CircleProgress.vue'
import ViewLayout from '@/components/ViewLayout.vue'
import StopButton from '@/components/StopButton.vue'

</script>

<template>
  <ViewLayout>
    <template #heading>Productivity</template>
    <CircleProgress />
    <template #after>
      <StopButton />
    </template>
  </ViewLayout>
</template>
