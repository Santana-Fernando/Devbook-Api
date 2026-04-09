package repositorios

import (
	"api/src/modelos"
	"errors"

	"gorm.io/gorm"
)

type Usuarios struct {
	db *gorm.DB
}

func NovoRepositorioDeUsuarios(db *gorm.DB) *Usuarios {
	return &Usuarios{db}
}

func (u Usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
	result := u.db.Create(&usuario)

	if result.Error != nil {
		return 0, result.Error
	}

	return uint64(usuario.ID), nil
}

func (u Usuarios) Buscar(nomeOuNick string) ([]modelos.Usuario, error) {
	var usuarios []modelos.Usuario

	resultado := u.db.
		Omit("senha").
		Where("nome LIKE ? OR nick LIKE ?", "%"+nomeOuNick+"%", "%"+nomeOuNick+"%").
		Find(&usuarios)

	if resultado.Error != nil {
		return nil, resultado.Error
	}

	return usuarios, nil
}

func (u Usuarios) BuscarPeloId(id uint64) (modelos.Usuario, error) {
	var usuario modelos.Usuario

	resultado := u.db.
		Omit("senha").
		Where("id = ?", id).
		Find(&usuario)

	if resultado.Error != nil {
		return modelos.Usuario{}, resultado.Error
	}

	return usuario, nil
}

func (u Usuarios) Atualizar(id uint64, usuario modelos.Usuario) error {
	resultado := u.db.Model(&modelos.Usuario{}).Where("id = ?", id).Updates(usuario)

	if resultado.Error != nil {
		return resultado.Error
	}

	return nil
}

func (u Usuarios) Deletar(id uint64) error {
	resultado := u.db.Where("id = ?", id).Delete(&modelos.Usuario{})

	if resultado.Error != nil {
		return resultado.Error
	}

	return nil
}

func (u Usuarios) BuscarPorEmail(email string) (modelos.Usuario, error) {
	var usuario modelos.Usuario

	resultado := u.db.
		Select("id, senha").
		Where("email = ?", email).
		First(&usuario)

	if resultado.Error != nil {
		return modelos.Usuario{}, resultado.Error
	}

	return usuario, nil
}

func (u Usuarios) Seguir(usuarioID, seguidorID uint64) error {
	if usuarioID == seguidorID {
		return errors.New("usuário não pode seguir a si mesmo")
	}

	var relacao modelos.Seguidores

	// Verifica se já existe
	err := u.db.
		Where("usuario_id = ? AND seguidor_id = ?", usuarioID, seguidorID).
		First(&relacao).Error

	// Se NÃO encontrou, pode inserir
	if errors.Is(err, gorm.ErrRecordNotFound) {
		novaRelacao := modelos.Seguidores{
			UsuarioID:  usuarioID,
			SeguidorID: seguidorID,
		}

		return u.db.Create(&novaRelacao).Error
	}

	// Se encontrou, não faz nada
	if err == nil {
		return nil // já segue, não insere novamente
	}

	// Se deu outro erro
	return err
}

func (u Usuarios) PararDeSeguir(usuarioID, seguidorID uint64) error {
	if usuarioID == seguidorID {
		return errors.New("usuário não pode deixar de seguir a si mesmo")
	}

	return u.db.
		Where("usuario_id = ? AND seguidor_id = ?", usuarioID, seguidorID).
		Delete(&modelos.Seguidores{}).Error
}

func (u Usuarios) BuscarSeguidores(usuarioID uint64) ([]modelos.Usuario, error) {
	var seguidores []modelos.Usuario

	resultado := u.db.
		Table("usuarios").
		Select("usuarios.id, usuarios.nome, usuarios.nick, usuarios.email, usuarios.criado_em").
		Joins("JOIN seguidores ON seguidores.seguidor_id = usuarios.id").
		Where("seguidores.usuario_id = ?", usuarioID).
		Find(&seguidores)

	if resultado.Error != nil {
		return nil, resultado.Error
	}

	return seguidores, nil
}

func (u Usuarios) BuscarSeguindo(usuarioID uint64) ([]modelos.Usuario, error) {
	var usuariosSeguidos []modelos.Usuario

	resultado := u.db.
		Table("usuarios").
		Select("usuarios.id, usuarios.nome, usuarios.nick, usuarios.email, usuarios.criado_em").
		Joins("JOIN seguidores ON seguidores.usuario_id = usuarios.id").
		Where("seguidores.seguidor_id = ?", usuarioID).
		Find(&usuariosSeguidos)

	if resultado.Error != nil {
		return nil, resultado.Error
	}

	return usuariosSeguidos, nil
}

func (u Usuarios) BuscarSenha(usuarioID uint64) (string, error) {
	var usuario modelos.Usuario

	resultado := u.db.
		Select("senha").
		Where("id = ?", usuarioID).
		First(&usuario)

	if resultado.Error != nil {
		return "", resultado.Error
	}

	return usuario.Senha, nil
}

func (u Usuarios) AtualizarSenha(usuarioID uint64, novaSenha string) error {
	resultado := u.db.Model(&modelos.Usuario{}).Where("id = ?", usuarioID).Update("senha", novaSenha)

	if resultado.Error != nil {
		return resultado.Error
	}

	return nil
}
