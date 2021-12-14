<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use App\Models\Cliente;

class ClienteController extends Controller
{
    //
    public function updateClient(Request $request, $id)
    {
        $client = Cliente::find($id);
        $client->nombre = $request->nombre;
        $client->ape_paterno = $request->ape_paterno;
        $client->ape_materno = $request->ape_materno;
        $client->correo_electronico = $request->correo_electronico;

        $client->save();

        return array(
            'ok' => true,
            'cliente' => $client
        );
    }
}
