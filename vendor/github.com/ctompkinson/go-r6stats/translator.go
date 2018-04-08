package r6

type translation struct {
	Raw   string
	Clean string
}

func translateOperatorSpecials(operatorName string, specials map[string]string) (clean map[string]string) {
	t := translations()
	opTranslations := t[operatorName]

	clean = map[string]string{}
	for _, t := range opTranslations {
		for k, v := range specials {
			if t.Raw == k {
				clean[t.Clean] = v
			}
		}
	}

	return
}

func translations() map[string][]translation {
	return map[string][]translation{
		"Hibana": {
			{
				Raw:   "operatorpvp_hibana_detonate_projectile",
				Clean: "Detonated Projectiles",
			},
		},
		"Vigil": {
			{
				Raw:   "operatorpvp_attackerdrone_diminishedrealitymode",
				Clean: "Diminished Reality Mode",
			},
		},
		"Lesion": {
			{
				Raw:   "operatorpvp_caltrop_enemy_affected",
				Clean: "Caltrop Enemy Affected",
			},
		},
		"Zofia": {
			{
				Raw:   "operatorpvp_concussiongrenade_detonate",
				Clean: "Concussion Grenades Detonated",
			},
		},
		"Fuze": {
			{
				Raw:   "operatorpvp_fuze_clusterchargekill",
				Clean: "Cluster Charge Kills",
			},
		},
		"Dokkaebi": {
			{
				Raw:   "operatorpvp_phoneshacked",
				Clean: "Phones Hacked",
			},
		},
		"Smoke": {
			{
				Raw:   "operatorpvp_smoke_poisongaskill",
				Clean: "Poison Gas Kills",
			},
		},
		"Finka": {
			{
				Raw:   "operatorpvp_rush_adrenalinerush",
				Clean: "Times Adrenaline Rushed",
			},
		},
		"Lion": {
			{
				Raw:   "operatorpvp_tagger_tagdevice_spot",
				Clean: "Tagging Drone Spots",
			},
		},
		"Doc": {
			{
				Raw:   "operatorpvp_doc_teammaterevive",
				Clean: "Teammate Revives",
			},
		},
		"Caveira": {
			{
				Raw:   "operatorpvp_caveira_interrogations",
				Clean: "Interrogations",
			},
		},
		"Sledge": {
			{
				Raw:   "operatorpvp_sledge_hammerkill",
				Clean: "Hammer Kills",
			},
		},
		"Castle": {
			{
				Raw:   "operatorpvp_castle_kevlarbarricadedeployed",
				Clean: "Kevlar Barricades Deployed",
			},
		},
		"Ela": {
			{
				Raw:   "operatorpvp_concussionmine_detonate",
				Clean: "Concussion Mines Detonated",
			},
		},
		"Glaz": {
			{
				Raw:   "operatorpvp_glaz_sniperpenetrationkill",
				Clean: "Sniper Penetration Kills",
			},
		},
		"Twitch": {
			{
				Raw:   "operatorpvp_twitch_gadgetdestroybyshockdrone",
				Clean: "Gadgets Destroyed By Shock Drone",
			},
		},
		"Tachanka": {
			{
				Raw:   "operatorpvp_tachanka_turretdeployed",
				Clean: "Turrets Deployed",
			},
		},
		"Thermite": {
			{
				Raw:   "operatorpvp_thermite_reinforcementbreached",
				Clean: "Reinforcements Breached",
			},
		},
		"Blitz": {
			{
				Raw:   "operatorpvp_blitz_flashfollowupkills",
				Clean: "Flash Follow Up Kills",
			},
		},
		"Ying": {
			{
				Raw:   "operatorpvp_dazzler_gadget_detonate",
				Clean: "Dazzler Gadget Detonated",
			},
		},
		"Mira": {
			{
				Raw:   "operatorpvp_black_mirror_gadget_deployed",
				Clean: "Black Mirrors Deployed",
			},
		},
		"Bandit": {
			{
				Raw:   "operatorpvp_bandit_batterykill",
				Clean: "Battery Kills",
			},
		},
		"Jackal": {
			{
				Raw:   "operatorpvp_cazador_assist_kill",
				Clean: "Cazador Assist Kills",
			},
		},
		"Echo": {
			{
				Raw:   "operatorpvp_echo_enemy_sonicburst_affected",
				Clean: "Echo Enemies Affected By Sonic Burst",
			},
		},
		"Pulse": {
			{
				Raw:   "operatorpvp_pulse_heartbeatassist",
				Clean: "Heart Beat Assists",
			},
		},
		"Montagne": {
			{
				Raw:   "operatorpvp_montagne_shieldblockdamage",
				Clean: "Shield Blocked Damage",
			},
		},
		"IQ": {
			{
				Raw:   "operatorpvp_iq_gadgetspotbyef",
				Clean: "Gadgets Spotted by EF",
			},
		},
		"Mute": {
			{
				Raw:   "operatorpvp_mute_jammerdeployed",
				Clean: "Jammers Deployed",
			},
		},
		"Ash": {
			{
				Raw:   "operatorpvp_ash_bonfirewallbreached",
				Clean: "Walls Breached By Bonfire",
			},
		},
		"Rook": {
			{
				Raw:   "operatorpvp_rook_armortakenteammate",
				Clean: "Armor Taken By Team Mates",
			},
		},
		"Capitão": {
			{
				Raw:   "operatorpvp_capitao_lethaldartkills",
				Clean: "Leathal Dart Kills",
			},
		},
		"Blackbeard": {
			{
				Raw:   "operatorpvp_blackbeard_gunshieldblockdamage",
				Clean: "Gun Shield Blocked Damage",
			},
		},
		"Buck": {
			{
				Raw:   "operatorpvp_buck_kill",
				Clean: "Master Key Kills",
			},
		},
		"Frost": {
			{
				Raw:   "operatorpvp_frost_dbno",
				Clean: "Frost DBNOs Placed",
			},
		},
		"Jäger": {
			{
				Raw:   "operatorpvp_jager_gadgetdestroybycatcher",
				Clean: "Gadgets Destroyed By Catcher",
			},
		},
		"Valkyrie": {
			{
				Raw:   "operatorpvp_valkyrie_camdeployed",
				Clean: "Cameras Deployed",
			},
		},
		"Thatcher": {
			{
				Raw:   "operatorpvp_thatcher_gadgetdestroywithemp",
				Clean: "Gadgets Destroyed With EMP",
			},
		},
		"Kapkan": {
			{
				Raw:   "operatorpvp_kapkan_boobytrapdeployed",
				Clean: "Booby Traps Placed",
			},
		},
	}
}
