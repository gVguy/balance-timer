<script lang="ts" setup>
import ViewLayout from '@/components/ViewLayout.vue'
import StopButton from '@/components/StopButton.vue'
import CircleProgress from '@/components/CircleProgress.vue'
import { SkipRest } from '@w/go/main/App'
import SkipButton from '@/components/SkipButton.vue'

</script>

<template>
  <ViewLayout>
    <template #heading>Rest</template>
    <CircleProgress />
    <span class="text">
      Take a break from the screen<br>
      and have some time for yourself.
    </span>
    <span class="text">
      Move around, look in the distance,<br>
      chat with someone, or just be mindful:<br>
      connect with reality and enjoy life.
    </span>
    <template #actions>
      <SkipButton
        title="Back to work"
        text="Press and hold to skip rest and immediately go to the next Productivity period"
        @confirm="SkipRest"
      />
      <StopButton />
    </template>
  </ViewLayout>
</template>
