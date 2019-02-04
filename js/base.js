$(document).ready(function() {
    console.log(document.cookie);
    console.log("Bienvenido a Inter Grafic");
    var formRegistro = $("#form_registro #cuerpo #submit_registro");
    console.log(formRegistro);

    $(formRegistro).click(function() {

        var nombre = $("#nombre").val();
        var user = $("#usuario_reg").val();
        var email = $("#email").val();
        var password = $("#contrasenia").val();

        console.log(nombre, user, email, password);        

        var datos = {
            nombre: nombre,
            usuario: user,
            email: email,
            password: password
        };

        console.log(datos);

        $.post({
            url:"/registro",
            data: JSON.stringify(datos),
            success: function(data, status, jqXHR) {
                console.log(data);
            },    

            dataType: "json"

        }).done(function(data) {
            console.log("Petición realizada");

        }).fail(function(data) {
            console.log("Petición fallida");        

        }).always(function(data){
            console.log("Petición completa");
        });

    });   
  

    // Ajax para iniciar sesion
    var formLogin = $(".waves-effect");
    console.log(formLogin)
    $(formLogin).click(function() {
        var usuario = $("#usuario").val();
        var password = $("#password").val();
        console.log(usuario, password);
        
        var login = {
            usuario: usuario,
            password: password
        };
        console.log(login);
        $.post({
            url:"/login",
            data: JSON.stringify(login),
            method: "POST",
            success: function(data, status, jqXHR) {
                console.log(data);
            },      
            dataType: "json"


        }).done(function(data) {
            console.log("Petición realizada");
            if (data == true){
                window.location.href = "/perfil";
            }
        
        }).fail(function(data) {
            console.log("Petición fallida");
            console.log(data);

        
        }).always(function(data){
            console.log("Petición completa");
        });
    });

});

