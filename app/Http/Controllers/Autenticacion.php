<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use Illuminate\Support\Facades\Hash;
use App\Models\Cliente;

class Autenticacion extends Controller
{
    public function store(Request $request)
    {

        // TODO(X ASER): Validar campos 
        $cliente = new Cliente();

        $cliente->nombre = $request->nombre;
        $cliente->ape_paterno = $request->ape_paterno;
        $cliente->ape_materno = $request->ape_materno;
        $cliente->contrasenia = Hash::make($request->contrasenia);
        $cliente->correo_electronico = $request->correo_electronico;

        $cliente->save();
        
        return array(
            "ok" => true,
            "cliente" => $cliente
        );
    }

    public function login(Request $request)
    {

        $cliente = Cliente::where('correo_electronico', $request->correo_electronico)
            ->first();

        
        if ($cliente && Hash::check($request->contrasenia, $cliente->contrasenia)) {
            return array(
                "ok" => true,
                "cliente" => $cliente
            );
        }
        
        return array(
            "ok" => false,
            "mensaje" => "Correo o contraseña inválidas"
        );
    }
}
