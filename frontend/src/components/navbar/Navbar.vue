<script setup lang="ts">
import UserMenu from "./UserMenu.vue"

var username = "username"
</script>


<!-- user menu haddling -->
<script lang=ts>
import { ref } from "vue"

var showUserMenu = ref(false)
function toggleShowUsermenu() {
    showUserMenu.value = !showUserMenu.value
}

var mouseLeaveMenu = ref(false)
function setMouseLeaveMenu(value:boolean) {
    mouseLeaveMenu.value = value
}

document.addEventListener("click", () => {
    if (mouseLeaveMenu.value) showUserMenu.value = false
    mouseLeaveMenu.value = false
})
</script>


<template>
    <div class="navbar">
        <div class="navbar-element" id="logo-btn">
            <a id="" href="{{}}"></a>
        </div>
        <nav class="navbar-element">
            <a class="nav-btn" href="/home">Home</a>
            <a class="nav-btn" href="/community">Comunidade</a>
        </nav>
        <div class="navbar-element" id="user-btn" @click="toggleShowUsermenu" @mouseleave="setMouseLeaveMenu(true)" @mouseenter="setMouseLeaveMenu(false)">
            <span id="username">{{ username }}</span>
            <img src="" alt="">
        </div>  
    </div>
    <UserMenu v-if="showUserMenu" @mouseleave="setMouseLeaveMenu(true)" @mouseenter="setMouseLeaveMenu(false)" />
</template>

<style>
.navbar {
    background-color: #ff0e6a;
    height: 70px;
    width: 100vw;
    position: fixed;
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 0 40px;
    box-sizing: border-box;
    user-select: none;
    color: white;
}

nav a {
    color: white;
    margin: 0 5px;
    text-decoration: none;
}

.navbar-element {
    cursor: pointer;
}
</style>