<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;

class Paso extends Model
{
    use HasFactory;
    protected $table = 'pasos';
    protected $fillable =
    [
        'id',
        'numero_orden',
        'contenido',
        'id_receta',
        'created_at',
        'updated_at'
    ];
}
