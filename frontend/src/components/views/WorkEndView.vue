<script lang="ts" setup>
import ViewLayout from '@/components/ViewLayout.vue'
import StopButton from '@/components/StopButton.vue'
import ContinueButton from '@/components/ContinueButton.vue'
import { useSession } from '@/composables/useSession'
import { computed } from 'vue'
import { BeginWork, BeginRest, MoreWork } from '@w/go/main/App'
import SkipButton from '@/components/SkipButton.vue'
import ExtendPeriodButton from '../ExtendPeriodButton.vue'

const { session } = useSession()

const restTime = computed(() => {
  const { restMin } = session.value.config
  const s = restMin == 1 ? '' : 's'
  return `${restMin} minute${s}`
})

</script>

<template>
  <ViewLayout>
    <template #heading>Have some rest</template>
    <ContinueButton @pointerdown="BeginRest" />
    <span class="countdown">Auto-continue to Rest in {{ session.timeRemainingString }}</span>
    <span>
      What you were doing is really important,<br>
      but it isn't going anywhere.
    </span>
    <span>
      You can come back to it in <span class="highlight">{{ restTime }}</span>.<br>
      For now, have some time for yourself.
    </span>
    <template #actions>
      <ExtendPeriodButton
        title="One more minute"
        text="Press and hold to go back to Productivity period for just one more minute"
        @confirm="MoreWork(1)"
      />
      <SkipButton
        title="Skip rest"
        text="Press and hold to move on to the next Productivity period without rest (not recommended)"
        @confirm="BeginWork"
      />
      <StopButton />
    </template>
  </ViewLayout>
</template>

<style scoped lang="scss">
.countdown {
  color: rgba(var(--fg-rgb), .3);
}
</style>
