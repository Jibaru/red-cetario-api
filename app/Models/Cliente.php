<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;

class Cliente extends Model
{
    use HasFactory;
    protected $table = 'clientes';
    protected $fillable =
    [
        'id',
        'nombre',
        'ape_paterno',
        'ape_materno',
        'contrasenia',
        'correo_electronico',
        'created_at',
        'updated_at'
    ];
}
