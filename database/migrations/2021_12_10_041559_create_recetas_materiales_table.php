<?php

use Illuminate\Database\Migrations\Migration;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Support\Facades\Schema;

class CreateRecetasMaterialesTable extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::create('recetas_materiales', function (Blueprint $table) {
            $table->unsignedBigInteger('id_receta');
            $table->foreign('id_receta')
                ->references('id')
                ->on('recetas')
                ->onDelete('cascade');
            $table->unsignedBigInteger('id_material');
            $table->foreign('id_material')
                ->references('id')
                ->on('materiales')
                ->onDelete('cascade');
        });
    }

    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down()
    {
        Schema::dropIfExists('recetas_materiales');
    }
}
