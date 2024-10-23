<script lang="ts" setup>
import CircleProgress from '@/components/CircleProgress.vue'
import ViewLayout from '@/components/ViewLayout.vue'
import StopButton from '@/components/StopButton.vue'
import { Skip } from '@w/go/main/App'
import SkipButton from '@/components/SkipButton.vue'

</script>

<template>
  <ViewLayout>
    <template #heading>Productivity</template>
    <CircleProgress />
    <span>
      Focus on your task.
    </span>
    <template #actions>
      <SkipButton
        title="Skip to rest"
        text="Press and hold to skip this Productivity period and begin the next Rest period"
        @confirm="Skip"
      />
      <StopButton />
    </template>
  </ViewLayout>
</template>
