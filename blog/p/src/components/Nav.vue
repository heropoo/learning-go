<template>
<!-- Start of nav bar  -->
<nav id="poppin-nav">
    <div id="nav-off"></div>
    <div id="header-logo-1" class="text-center">
        <a href="#"><img src="../assets/logo.png" alt="logo"></a>
    </div>
    <ul id="menu">
        <li class="has-submenu text-capitalize">
            <a href="index.html">home</a>
        </li>
        <li class="has-submenu">
            <a href="#">Blog <span class="ti-angle-down"></span></a>
            <ul class="submenu">
                <li><a href="blog.html">Blog Grid</a></li>
                <li><a href="blog-single.html">Single Post</a></li>
            </ul>
        </li>
        <li><a href="contact.html">Contact</a></li>
    </ul>
</nav>
<!-- <div id="menu-overlay"></div> -->
<!-- End of nav bar  -->
</template>

<script>
export default {
  name: 'Nav',
  props: {
    msg: String
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

</style>