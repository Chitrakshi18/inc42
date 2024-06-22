<?php

if ( ! defined( 'ABSPATH' ) ) {
    exit;
}

function myplugin_greet() {
    echo '<p>Hello from WordPress!</p>';
}

add_action( 'admin_notices', 'myplugin_greet' );
