<template>
    <div class="cursor-preview">
        <div class="custom-cursor">
            <div
                id="cursor-big"
                class="custom-cursor__ball custom-cursor__ball--big"></div>
            <div
                id="cursor-small"
                class="custom-cursor__ball custom-cursor__ball--small"></div>
        </div>
    </div>
</template>

<script lang="ts">
import { gsap } from "gsap";
export default {
    mounted() {
        const cursorBig = document.getElementById("cursor-big"),
            cursorSmall = document.getElementById("cursor-small"),
            links = document.getElementsByTagName("a"),
            withClassHover = document.getElementsByClassName("cursor-hover"),
            withHover = [...links, ...withClassHover];

        // Event Listeners
        document.addEventListener("mousemove", onMouseMove);
        document.addEventListener("mousedown", onMouseHover);
        document.addEventListener("mouseup", onMouseHoverOut);
        document.addEventListener("mouseenter", () => {
            cursorBig!.style.opacity = "1";
            cursorSmall!.style.opacity = "1";
        });
        document.addEventListener("mouseleave", () => {
            cursorBig!.style.opacity = "0";
            cursorSmall!.style.opacity = "0";
        });
        withHover.forEach((element) => {
            element.addEventListener("mouseover", onMouseHover);
            element.addEventListener("mouseout", onMouseHoverOut);
        });

        // Event Handlers
        function onMouseMove(e: any) {
            cursorSmall!.style.opacity = "1";
            gsap.to(cursorBig, 0.4, {
                x: e.clientX - 18,
                y: e.clientY - 18,
            });
            gsap.to(cursorSmall, 0.1, {
                x: e.clientX - 4,
                y: e.clientY - 4,
            });
        }
        function onMouseHover() {
            gsap.to(cursorBig, 0.3, {
                scale: 3,
            });
        }
        function onMouseHoverOut() {
            gsap.to(cursorBig, 0.3, {
                scale: 1,
            });
        }
    },
};
</script>

<style land="scss">
.cursor-preview__title {
    font-size: 20px;
    line-height: 24px;
    color: #8a9eff;
    margin-bottom: 5px;
}

.cursor-preview__link {
    color: #fff;
    font-weight: 600;
    font-size: 24px;
    line-height: 29px;
}

.cursor-preview__button {
    background: white;
    border-radius: 5px;
    padding: 16px 50px;
    color: #fff;
    font-weight: 600;
    font-size: 24px;
    line-height: 29px;
    outline: none;
    border: none;
}

.custom-cursor__ball {
    position: fixed;
    top: 0;
    left: 0;
    mix-blend-mode: difference;
    z-index: 99999;
    opacity: 0;
    pointer-events: none;
    transition: opacity 0.5s ease;
}

.custom-cursor__ball--big {
    content: "";
    width: 30px;
    height: 30px;
    background: white;
    border-radius: 50%;
}

.custom-cursor__ball--small {
    content: "";
    width: 6px;
    height: 6px;
    background: white;
    border-radius: 50%;
}
</style>
