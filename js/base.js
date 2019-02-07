$(document).ready(function() {
    ActualizarFotos();

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
            console.log(data);
            alert("Te has registrado con éxito");
            if (data == true){
                window.location.href = "/";
            } else {
                alert("El usuario ya existe.");
            }

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
                window.location.href = "/";
            }
        
        }).fail(function(data) {
            console.log("Petición fallida");
            console.log(data);

        
        }).always(function(data){
            console.log("Petición completa");
        });
   
    });

    //para ocultar los tres botones sin loguear o para ocultar el login en caso de iniciar sesión
    if (document.cookie != ""){
        $("#formLogin").hide();
    }else{
        $(".right2").hide();
    }

   

    
});


//Ajax para mostrar las fotos
function ActualizarFotos() {   
    $.ajax({
        url: "/listarFoto",
        method: "POST",
        dataType: "json",
        contentType: "application/json",
        success: function(data) {
            Historial_Fotos(data);
        },
        error: function(data) {
            console.log(data);
        }
    });
}

function Historial_Fotos(array) {
    
    var div = $("#fotos");
    div.children().remove();
    if(array != null && array.length > 0) {
      

        for(var x = 0; x < array.length; x++) {
            div.append( 
            "<div>"
                +"<img src='/files/"+array[x].URL+"' width='300px' height='180px'>"+
                "<p>"+array[x].Texto+"</p>"+
            "</div>");
        }
    } else {
        div.append('<div colspan="3">No hay registros de hoy</div>');
        
    }
}