<script lang="ts">
export default {
    data() {
        return {
            typeValue: "",
            typeStatus: false,
            displayText: {
                fullName: "Ivan Tomić",
                location: "Croatia",
                techStack: ["Java", "Spring Boot", "Docker", "PostgreSQL"],
            },
            typingSpeed: 20,
            charIndex: 0,
        };
    },
    mounted() {
        this.typeText();
    },
    methods: {
        typeText() {
            if (
                this.charIndex <
                JSON.stringify(this.displayText, null, 4).length
            ) {
                if (!this.typeStatus) this.typeStatus = true;
                this.typeValue += JSON.stringify(
                    this.displayText,
                    null,
                    4
                ).charAt(this.charIndex);
                this.charIndex += 1;
                requestAnimationFrame(this.typeText);
            } else {
                this.typeStatus = false;
            }
        },
    },
};
</script>

<template>
    <section
        class="flex justify-around items-center sm:mt-40 mt-32 lg:flex-row flex-col lg:gap-0 gap-10">
        <div>
            <p class="text-5xl sm:text-6xl font-semibold leading-12">Hi, I am</p>
            <h1 class="text-5xl sm:text-6xl font-bold leading-12">
                Ivan Tomić
            </h1>
            <p class="text-2xl">Software Developer</p>
        </div>
        <div
            class="h-80 w-80 bg-black rounded-lg p-4 flex flex-col justify-center">
            <div class="flex justify-end gap-2">
                <div class="h-3 w-3 bg-green rounded-full hover:bg-opacity-80" />
                <div class="h-3 w-3 bg-orange rounded-full hover:bg-opacity-80" />
                <div class="h-3 w-3 bg-red rounded-full hover:bg-opacity-80" />
            </div>
            <div class="h-full w-full flex justify-center items-center">
                <div class="h-full w-full flex items-center">
                    <pre class="text-white inline">{{ typeValue }}<p class="blinking-cursor text-white inline">|</p></pre>
                </div>
            </div>
        </div>
    </section>
</template>

<style>
.blinking-cursor {
    -webkit-animation: 1s blink step-end infinite;
    -moz-animation: 1s blink step-end infinite;
    -ms-animation: 1s blink step-end infinite;
    -o-animation: 1s blink step-end infinite;
    animation: 1s blink step-end infinite;
}
@keyframes blink {
    from,
    to {
        color: transparent;
    }
    50% {
        color: white;
    }
}
@-moz-keyframes blink {
    from,
    to {
        color: transparent;
    }
    50% {
        color: white;
    }
}
@-webkit-keyframes blink {
    from,
    to {
        color: transparent;
    }
    50% {
        color: white;
    }
}
@-ms-keyframes blink {
    from,
    to {
        color: transparent;
    }
    50% {
        color: white;
    }
}
@-o-keyframes blink {
    from,
    to {
        color: transparent;
    }
    50% {
        color: white;
    }
}
</style>
