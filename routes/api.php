<?php

use Illuminate\Http\Request;
use Illuminate\Support\Facades\Route;
use App\Http\Controllers\Autenticacion;
use App\Http\Controllers\RecetaController;
use App\Http\Controllers\NotificacionController;
use App\Http\Controllers\ClienteController;

/*
|--------------------------------------------------------------------------
| API Routes
|--------------------------------------------------------------------------
|
| Here is where you can register API routes for your application. These
| routes are loaded by the RouteServiceProvider within a group which
| is assigned the "api" middleware group. Enjoy building your API!
|
*/

Route::middleware('auth:sanctum')->get('/user', function (Request $request) {
    return $request->user();
});

Route::post('/clientes', [Autenticacion::class, 'store']);
Route::post('/login', [Autenticacion::class, 'login']);

Route::get('/recetas', [RecetaController::class, 'index']);
Route::get('/receta/{id}', [RecetaController::class, 'show']);
Route::post('/recetas/{id}/comentario', [RecetaController::class, 'comentar']);

Route::get('/notificaciones', [NotificacionController::class, 'index']);
Route::put('/notificaciones/{id}/fecha-visto', [NotificacionController::class, 'updateNotificacion']);
Route::delete('/notificacion/{id}', [NotificacionController::class, 'destroy']);
Route::delete('/notificaciones/cliente/{id}', [NotificacionController::class, 'destroyForClient']);

Route::put('/clientes/{id}', [ClienteController::class, 'updateClient']);